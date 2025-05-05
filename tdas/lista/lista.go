package listas

type Lista[T any] interface {

	//EstaVacia devuelve verdadero si la lista no tiene elementos listados, false en caso contrario.
	EstaVacia() bool

	//InsertarPrimero inserta un elemento al inicio de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta un elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero borra el primer elemento de la lista. Si la lista tiene elementos elimina y devuelve
	// el primer valor. Si la lista esta vacia entra en panico con el mensaje "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero obtiene el valor del primer elemento de la lista. Si la lista esta vacia
	// entra en panico con el mensaje "La lista esta vacia"
	VerPrimero() T

	//VerPrimero obtiene el valor del ultimo elemento de la lista. Si la lista esta vacia
	// entra en panico con el mensaje "La lista esta vacia"
	VerUltimo() T

	//Largo devuelve la cantidad de elementos guardados en una lista
	Largo() int

	//Iterar recorre todos los elementos de la lista, ejecutando la función 'visitar' en cada uno.
	//La función 'visitar' recibe el elemento y debe devolver true para continuar o false para detener la iteración.
	Iterar(visitar func(T) bool)

	//Iterador es un iterador externo de la lista.
	//Iterador crea y devuelve un iterador externo para recorrer la lista elemento por elemento manualmente.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//VerActual devuelve el valor sobre el que está parado el iterador. En caso de haber iterador todos
	// los elementos, entra en panico con el mensaje "El iterador termino de iterar"
	VerActual() T

	//HaySiguiente devuelve verdadero si hay un valor valido en la posición actual o falso en caso contrario.
	HaySiguiente() bool

	//Siguiente mueve el iterador al siguiente nodo. En caso de que no haya siguiente,
	//  entra en panico con el mensaje "El iterador termino de iterar"
	Siguiente()

	//Insertar inserta el valor T en la posición actual
	Insertar(T)

	//Borrar borra y devuelve el valor actual, pasando el iterador al valor siguiente. En caso de haber
	// iterado todos los elementos, entra en panico con el mensaje "El iterador termino de iterar"
	Borrar() T
}
