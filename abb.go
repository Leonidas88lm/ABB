package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoABB[K comparable, V any] struct {
	clave K
	dato  V
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		cmp:      cmp,
	}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.guardarRec(a.raiz, clave, dato)
}

func (a *abb[K, V]) guardarRec(n *nodoABB[K, V], clave K, dato V) *nodoABB[K, V] {
	if n == nil {
		a.cantidad++
		return &nodoABB[K, V]{clave: clave, dato: dato}
	}
	cmp := a.cmp(clave, n.clave)
	if cmp < 0 {
		n.izq = a.guardarRec(n.izq, clave, dato)
	} else if cmp > 0 {
		n.der = a.guardarRec(n.der, clave, dato)
	} else {
		n.dato = dato
	}
	return n
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscarNodo(a.raiz, clave) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo := a.buscarNodo(a.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (a *abb[K, V]) buscarNodo(n *nodoABB[K, V], clave K) *nodoABB[K, V] {
	if n == nil {
		return nil
	}
	cmp := a.cmp(clave, n.clave)
	if cmp < 0 {
		return a.buscarNodo(n.izq, clave)
	} else if cmp > 0 {
		return a.buscarNodo(n.der, clave)
	}
	return n
}

func (a *abb[K, V]) Borrar(clave K) V {
	var borrado V
	var ok bool
	a.raiz, borrado, ok = a.borrarRec(a.raiz, clave)
	if !ok {
		panic("La clave no pertenece al diccionario")
	}
	a.cantidad--
	return borrado
}

func (a *abb[K, V]) borrarRec(n *nodoABB[K, V], clave K) (*nodoABB[K, V], V, bool) {
	if n == nil {
		var cero V
		return nil, cero, false
	}
	cmp := a.cmp(clave, n.clave)
	if cmp < 0 {
		var borrado V
		var ok bool
		n.izq, borrado, ok = a.borrarRec(n.izq, clave)
		return n, borrado, ok
	}
	if cmp > 0 {
		var borrado V
		var ok bool
		n.der, borrado, ok = a.borrarRec(n.der, clave)
		return n, borrado, ok
	}

	// Caso encontrado
	borrado := n.dato
	if n.izq == nil && n.der == nil {
		return nil, borrado, true
	}
	if n.izq == nil {
		return n.der, borrado, true
	}
	if n.der == nil {
		return n.izq, borrado, true
	}

	// Caso con dos hijos: reemplazar por sucesor in-order
	sucesor := a.buscarMin(n.der)
	n.clave = sucesor.clave
	n.dato = sucesor.dato
	n.der, _, _ = a.borrarRec(n.der, sucesor.clave)
	return n, borrado, true
}

func (a *abb[K, V]) buscarMin(n *nodoABB[K, V]) *nodoABB[K, V] {
	for n.izq != nil {
		n = n.izq
	}
	return n
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.iterarInOrder(a.raiz, visitar)
}

func (a *abb[K, V]) iterarInOrder(n *nodoABB[K, V], visitar func(K, V) bool) bool {
	if n == nil {
		return true
	}
	if !a.iterarInOrder(n.izq, visitar) {
		return false
	}
	if !visitar(n.clave, n.dato) {
		return false
	}
	return a.iterarInOrder(n.der, visitar)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	a.iterarRango(a.raiz, desde, hasta, visitar)
}

func (a *abb[K, V]) iterarRango(n *nodoABB[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {
	if n == nil {
		return true
	}
	if desde == nil || a.cmp(n.clave, *desde) >= 0 {
		if !a.iterarRango(n.izq, desde, hasta, visitar) {
			return false
		}
	}
	if (desde == nil || a.cmp(n.clave, *desde) >= 0) &&
		(hasta == nil || a.cmp(n.clave, *hasta) <= 0) {
		if !visitar(n.clave, n.dato) {
			return false
		}
	}
	if hasta == nil || a.cmp(n.clave, *hasta) <= 0 {
		return a.iterarRango(n.der, desde, hasta, visitar)
	}
	return true
}

// iteradorABB es una estructura auxiliar para implementar el iterador, utilizando una pila como estructura auxiliar
type iteradorABB[K comparable, V any] struct {
	pila   TDAPila.Pila[*nodoABB[K, V]]
	actual *nodoABB[K, V]
	cmp    func(K, K) int
	desde  *K
	hasta  *K
}

// apilarDesdeHasta apila todos los hijos izquierdos del nodo que recibe, que se encuentren en el rango [desde,hasta].
//
//	En caso que un limite sea nil, no lo tiene en cuenta
func (it *iteradorABB[K, V]) apilarDesdeHasta(nodo *nodoABB[K, V], desde *K, hasta *K) {
	for nodo != nil {

		if desde != nil && it.cmp(nodo.clave, *desde) < 0 {
			nodo = nodo.der
		} else if hasta != nil && it.cmp(nodo.clave, *hasta) > 0 {
			nodo = nodo.izq
		} else {
			it.pila.Apilar(nodo)
			nodo = nodo.izq
		}
	}
}

func (iterABB *iteradorABB[K, V]) panicIterABB() {
	if iterABB.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iter := &iteradorABB[K, V]{pila: pila, cmp: abb.cmp, desde: nil, hasta: nil}
	iter.apilarDesdeHasta(abb.raiz, iter.desde, iter.hasta)
	return iter
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	iter := &iteradorABB[K, V]{pila: pila, cmp: abb.cmp, desde: desde, hasta: hasta}
	iter.apilarDesdeHasta(abb.raiz, iter.desde, iter.hasta)
	return iter
}

func (iterABB *iteradorABB[K, V]) HaySiguiente() bool {
	return !iterABB.pila.EstaVacia()
}

func (iterABB *iteradorABB[K, V]) VerActual() (K, V) {
	iterABB.panicIterABB()
	nodoActual := iterABB.pila.VerTope()
	return nodoActual.clave, nodoActual.dato
}

func (iterABB *iteradorABB[K, V]) Siguiente() {
	iterABB.panicIterABB()
	nodoActual := iterABB.pila.Desapilar()
	if iterABB.actual.der != nil {
		iterABB.apilarDesdeHasta(nodoActual.der, iterABB.desde, iterABB.hasta)
	}
}
