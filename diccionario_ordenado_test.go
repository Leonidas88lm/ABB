package diccionario_test

import (
	"cmp"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](cmp.Compare)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestDiccionarioOrdenadoUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar(claves[0], valores[0])

	iter := dic.Iterador()
	keys := []string{}
	values := []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}

	require.EqualValues(t, []string{claves[0]}, keys)
	require.EqualValues(t, []string{valores[0]}, values)

	dic.Guardar(claves[1], valores[1])
	iter = dic.Iterador()
	keys = []string{}
	values = []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}

	require.EqualValues(t, []string{claves[0], claves[1]}, keys)
	require.EqualValues(t, []string{valores[0], valores[1]}, values)

	dic.Guardar(claves[2], valores[2])
	iter = dic.Iterador()
	keys = []string{}
	values = []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}

	require.EqualValues(t, []string{claves[0], claves[1], claves[2]}, keys)
	require.EqualValues(t, []string{valores[0], valores[1], valores[2]}, values)

}

func TestDiccionarioOrdenadoDerecha(t *testing.T) {
	t.Log("Prueba con todos los elementos agregados al lado derecho del árbol")

	claves := []string{"A", "B", "C", "D", "E"}
	valores := []string{"a", "b", "c", "d", "e"}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	iter := dic.Iterador()
	keys := []string{}
	values := []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}

	require.EqualValues(t, claves, keys)
	require.EqualValues(t, valores, values)
}

func TestDiccionarioOrdenadoIzquierda(t *testing.T) {
	t.Log("Prueba con todos los elementos agregados al lado derecho del árbol")

	claves := []string{"E", "D", "C", "B", "A"}
	valores := []string{"e", "d", "c", "b", "a"}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	iter := dic.Iterador()
	keys := []string{}
	values := []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}

	require.EqualValues(t, []string{"A", "B", "C", "D", "E"}, keys)
	require.EqualValues(t, []string{"a", "b", "c", "d", "e"}, values)
}

func TestDiccionarioOrdenadoVariado(t *testing.T) {
	t.Log("Prueba con datos insertados en orden aleatorio")

	claves := []string{"Gato", "Perro", "Vaca", "Pajaro", "Cerdo"}
	valores := []string{"miau", "guau", "moo", "pio", "oink"}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	iter := dic.Iterador()
	keys := []string{}
	values := []string{}
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		keys = append(keys, clave)
		values = append(values, valor)
		iter.Siguiente()
	}
	//ElementsMatch se fija que todos los valores esten aunque no importa el orden
	require.ElementsMatch(t, claves, keys)
	require.ElementsMatch(t, valores, values)
}

func TestDiccionarioOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestDiccionarioOrdenadoConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestDiccionarioOridenadoConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}
	compararStruct := func(a, b avanzado) int {
		if res := cmp.Compare(a.w, b.w); res != 0 {
			return res
		}
		if res := strings.Compare(a.z, b.z); res != 0 {
			return res
		}
		if res := strings.Compare(a.x.a, b.x.a); res != 0 {
			return res
		}
		if res := cmp.Compare(a.x.b, b.x.b); res != 0 {
			return res
		}
		if res := strings.Compare(a.y.a, b.y.a); res != 0 {
			return res
		}
		return cmp.Compare(a.y.b, b.y.b)
	}

	dic := TDADiccionario.CrearABB[avanzado, int](compararStruct)

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestDccionarioOrdenadoClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestDiccionarioOrdenadoValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}
