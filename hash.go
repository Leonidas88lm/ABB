package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const CAPACIDAD_INICIAL = 13
const FACTOR_CARGA_MAXIMO = 0.75
const FACTOR_CARGA_MINIMO = 0.25
const FACTOR_REDIMENSION = 2

type entrada[K comparable, V any] struct {
	clave K
	valor V
}

type diccionarioHash[K comparable, V any] struct {
	capacidad int
	tabla     []TDALista.Lista[entrada[K, V]]
	cantidad  int
}

// convertirABytes convierte a un array de bytes la clave recibida
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// calcularIndice hace un hash con la clave para generar un indice
func (diccionario *diccionarioHash[K, V]) calcularIndice(clave K, capacidad int) int {
	bytes := convertirABytes(clave)
	var hash uint32 = 0
	for _, b := range bytes {
		hash = hash*31 + uint32(b)
	}
	return int(hash % uint32(capacidad))
}

// crearIteradorPosicionado genera un índice con la clave, crea un iterador sobre la lista correspondiente,
// y lo posiciona en la clave si existe, o al final de la lista si no se encuentra.
func (diccionario *diccionarioHash[K, V]) crearIteradorPosicionado(clave K) TDALista.IteradorLista[entrada[K, V]] {
	indice := diccionario.calcularIndice(clave, diccionario.capacidad)
	lista := diccionario.tabla[indice]
	if lista == nil {
		lista = TDALista.CrearListaEnlazada[entrada[K, V]]()
		diccionario.tabla[indice] = lista
	}
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			break
		}
		iter.Siguiente()
	}
	return iter
}

// panicDiccionario devuelve un panic si la lista esta vacia
func (diccionario *diccionarioHash[K, V]) panicDiccionario(clave K) {
	if !diccionario.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
}

// deboAumentar verifica si es necesario incrementar el tamaño del diccionario
func (diccionario *diccionarioHash[K, V]) deboAumentar(cantidad int, capacidad int) bool {
	factorCarga := float32(cantidad) / float32(capacidad)
	return factorCarga >= FACTOR_CARGA_MAXIMO
}

// debodisminuir verifica si es necesario reducir el tamaño del diccionario
func (diccionario *diccionarioHash[K, V]) deboDisminuir(cantidad int, capacidad int) bool {
	factorCarga := float32(cantidad) / float32(capacidad)
	return (factorCarga <= FACTOR_CARGA_MINIMO)
}

// rehash moficia el tamaño del diccionario, actualizando cada par clave - valor
func (diccionario *diccionarioHash[K, V]) rehash(nuevaCapacidad int) {
	nuevaTabla := make([]TDALista.Lista[entrada[K, V]], nuevaCapacidad)
	for i := 0; i < diccionario.capacidad; i++ {
		lista := diccionario.tabla[i]
		if lista != nil {
			iter := lista.Iterador()
			for iter.HaySiguiente() {
				elementoActual := iter.VerActual()
				indice := diccionario.calcularIndice(elementoActual.clave, nuevaCapacidad)
				if nuevaTabla[indice] == nil {
					nuevaTabla[indice] = TDALista.CrearListaEnlazada[entrada[K, V]]()
				}
				nuevaTabla[indice].InsertarUltimo(elementoActual)
				iter.Siguiente()
			}
		}
	}
	diccionario.tabla = nuevaTabla
	diccionario.capacidad = nuevaCapacidad
}

func (diccionario *diccionarioHash[K, V]) Guardar(clave K, dato V) {

	if diccionario.deboAumentar(diccionario.cantidad, diccionario.capacidad) {
		diccionario.rehash(diccionario.capacidad * FACTOR_REDIMENSION)
	}

	iter := diccionario.crearIteradorPosicionado(clave)
	indice := diccionario.calcularIndice(clave, diccionario.capacidad)

	if iter.HaySiguiente() {
		iter.Borrar()
	} else {
		diccionario.cantidad++
	}

	diccionario.tabla[indice].InsertarUltimo(entrada[K, V]{clave, dato})
}

func (diccionario *diccionarioHash[K, V]) Pertenece(clave K) bool {
	iter := diccionario.crearIteradorPosicionado(clave)
	return iter.HaySiguiente()
}

func (diccionario *diccionarioHash[K, V]) Obtener(clave K) V {
	diccionario.panicDiccionario(clave)
	iter := diccionario.crearIteradorPosicionado(clave)
	return iter.VerActual().valor
}

func (diccionario *diccionarioHash[K, V]) Borrar(clave K) V {
	diccionario.panicDiccionario(clave)
	if diccionario.deboDisminuir(diccionario.cantidad, diccionario.capacidad) {
		nuevaCapacidad := max(diccionario.capacidad/FACTOR_REDIMENSION, CAPACIDAD_INICIAL)
		diccionario.rehash(nuevaCapacidad)
	}
	iter := diccionario.crearIteradorPosicionado(clave)
	valorBorrado := iter.Borrar().valor
	diccionario.cantidad--
	return valorBorrado
}

func (diccionario *diccionarioHash[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[entrada[K, V]], CAPACIDAD_INICIAL)
	return &diccionarioHash[K, V]{
		capacidad: CAPACIDAD_INICIAL,
		tabla:     tabla,
		cantidad:  0,
	}
}
