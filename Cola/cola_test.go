package cola_test

import (
	TDACola "cola"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de una cantidad baja de elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, true, cola.EstaVacia())
}

func TestPilaPocosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, true, cola.EstaVacia())
	for i := 0; i <= 3; i++ {
		cola.Encolar(i)
		require.EqualValues(t, false, cola.EstaVacia())
		require.EqualValues(t, 0, cola.VerPrimero())
		cola.Encolar(i)
		require.EqualValues(t, false, cola.EstaVacia())
		require.EqualValues(t, 0, cola.VerPrimero())
		cola.Encolar(i)
		require.EqualValues(t, false, cola.EstaVacia())
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for i := 0; i <= 3; i++ {
		cola.Desencolar()
		cola.Desencolar()
		cola.Desencolar()
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, true, cola.EstaVacia())

}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos la prueba de volumen")
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i <= 10000; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
		require.EqualValues(t, false, cola.EstaVacia())
	}
	for i := 0; i <= 10000; i++ {
		require.EqualValues(t, false, cola.EstaVacia())
		require.EqualValues(t, i, cola.VerPrimero())
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.EqualValues(t, true, cola.EstaVacia())

	colaStr := TDACola.CrearColaEnlazada[string]()
	var alfabeto [26]string = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, letra := range alfabeto {
		colaStr.Encolar(letra)
		require.EqualValues(t, "A", colaStr.VerPrimero())
		require.EqualValues(t, false, colaStr.EstaVacia())
	}
	for i := 0; i < len(alfabeto)-1; i++ {
		require.EqualValues(t, alfabeto[i], colaStr.Desencolar())
		require.EqualValues(t, alfabeto[i+1], colaStr.VerPrimero())
		require.EqualValues(t, false, colaStr.EstaVacia())
	}
	require.EqualValues(t, alfabeto[25], colaStr.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.EqualValues(t, true, colaStr.EstaVacia())
}
