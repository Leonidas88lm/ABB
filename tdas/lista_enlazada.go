package lista

type nodoLista[T any] struct {
	dato        T
	proximoNodo *nodoLista[T]
}

// Crea un nodo nuevo
func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{
		dato:        dato,
		proximoNodo: nil,
	}

}

// panicLista devuelve un panic si la lista esta vacia
func (lista *listaEnlazada[T]) panicLista() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)
	nuevoNodo.proximoNodo = lista.primero
	lista.primero = nuevoNodo
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		lista.ultimo.proximoNodo = nuevoNodo
		lista.ultimo = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.panicLista()
	dato := lista.primero.dato
	lista.primero = lista.primero.proximoNodo
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.panicLista()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.panicLista()
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterar aplica la función visitar a cada elemento de la lista
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero
	for nodoActual != nil {
		if !visitar(nodoActual.dato) {
			break
		}
		nodoActual = nodoActual.proximoNodo
	}
}

// Iterador devuelve un iterador externo para la lista
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {

	return &iteradorListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// panicIterador devuelve un panic si el iterador ya paso la ultima posicion
func (it *iteradorListaEnlazada[T]) panicIterador() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (it *iteradorListaEnlazada[T]) VerActual() T {
	it.panicIterador()
	return it.actual.dato
}

func (it iteradorListaEnlazada[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iteradorListaEnlazada[T]) Siguiente() {
	it.panicIterador()
	it.anterior = it.actual
	it.actual = it.actual.proximoNodo
}

func (it *iteradorListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	if it.anterior == nil {
		if it.lista.primero == nil {
			it.lista.primero = nuevoNodo
			it.lista.ultimo = nuevoNodo
		} else {
			nuevoNodo.proximoNodo = it.lista.primero
			it.lista.primero = nuevoNodo
		}
	} else if it.actual == nil {
		it.anterior.proximoNodo = nuevoNodo
		it.lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.proximoNodo = it.actual
		it.anterior.proximoNodo = nuevoNodo
	}

	it.actual = nuevoNodo
	it.lista.largo++
}

func (it *iteradorListaEnlazada[T]) Borrar() T {
	it.panicIterador()

	datoBorrado := it.actual.dato

	if it.anterior == nil {
		it.lista.primero = it.actual.proximoNodo
	} else {
		it.anterior.proximoNodo = it.actual.proximoNodo
	}

	if it.actual == it.lista.ultimo {
		it.lista.ultimo = it.anterior
	}

	it.actual = it.actual.proximoNodo
	it.lista.largo--

	return datoBorrado
}
