package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de una cantidad baja de elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, true, pila.EstaVacia())
	t.Log("Veo si la pila vacia se sigue comportando como tal incluso despues de apilar y desapilar")
	pila.Apilar(1)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, true, pila.EstaVacia())
}

func TestLIFO(t *testing.T) {
	t.Log("Pruebo que los elementos son devueltos en el orden correcto al desapilarlos")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Apilar(6)
	pila.Apilar(7)
	pila.Apilar(8)
	pila.Apilar(9)
	require.EqualValues(t, 9, pila.Desapilar())
	require.EqualValues(t, 8, pila.Desapilar())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 6, pila.Desapilar())
	require.EqualValues(t, 5, pila.Desapilar())
	t.Log("Apilo un item en el medio, por las dudas")
	pila.Apilar(99)
	require.EqualValues(t, 99, pila.Desapilar())
	t.Log("Miro el tope en el medio")
	require.EqualValues(t, 4, pila.VerTope())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.EqualValues(t, true, pila.EstaVacia())
	t.Log("Pruebo que los elementos son devueltos en el orden correcto al desapilarlos, esta vez con strings")
	pilaStr := TDAPila.CrearPilaDinamica[string]()
	pilaStr.Apilar("A")
	pilaStr.Apilar("B")
	pilaStr.Apilar("C")
	pilaStr.Apilar("D")
	pilaStr.Apilar("E")
	pilaStr.Apilar("F")
	pilaStr.Apilar("G")
	pilaStr.Apilar("H")
	pilaStr.Apilar("I")
	pilaStr.Apilar("J")
	pilaStr.Apilar("K")
	pilaStr.Apilar("L")
	pilaStr.Apilar("M")
	pilaStr.Apilar("N")
	pilaStr.Apilar("O")
	pilaStr.Apilar("P")
	pilaStr.Apilar("Q")
	pilaStr.Apilar("R")
	pilaStr.Apilar("S")
	pilaStr.Apilar("T")
	pilaStr.Apilar("U")
	pilaStr.Apilar("V")
	pilaStr.Apilar("W")
	pilaStr.Apilar("X")
	pilaStr.Apilar("Y")
	pilaStr.Apilar("Z")
	require.EqualValues(t, "Z", pilaStr.VerTope())
	require.EqualValues(t, "Z", pilaStr.Desapilar())
	require.EqualValues(t, "Y", pilaStr.Desapilar())
	require.EqualValues(t, "X", pilaStr.Desapilar())
	require.EqualValues(t, "W", pilaStr.Desapilar())
	require.EqualValues(t, "V", pilaStr.Desapilar())
	require.EqualValues(t, "U", pilaStr.Desapilar())
	require.EqualValues(t, "T", pilaStr.Desapilar())
	require.EqualValues(t, "S", pilaStr.Desapilar())
	require.EqualValues(t, "R", pilaStr.Desapilar())
	require.EqualValues(t, "Q", pilaStr.Desapilar())
	require.EqualValues(t, "P", pilaStr.Desapilar())
	require.EqualValues(t, "O", pilaStr.Desapilar())
	require.EqualValues(t, "N", pilaStr.Desapilar())
	require.EqualValues(t, "M", pilaStr.Desapilar())
	require.EqualValues(t, "L", pilaStr.Desapilar())
	require.EqualValues(t, "K", pilaStr.Desapilar())
	require.EqualValues(t, "J", pilaStr.Desapilar())
	require.EqualValues(t, "I", pilaStr.Desapilar())
	require.EqualValues(t, "H", pilaStr.Desapilar())
	require.EqualValues(t, "G", pilaStr.Desapilar())
	require.EqualValues(t, "F", pilaStr.Desapilar())
	require.EqualValues(t, "E", pilaStr.Desapilar())
	require.EqualValues(t, "D", pilaStr.Desapilar())
	require.EqualValues(t, "C", pilaStr.Desapilar())
	require.EqualValues(t, "B", pilaStr.Desapilar())
	require.EqualValues(t, "A", pilaStr.Desapilar())
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos la prueba de volumen")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 10000; i++ {
		pila.Apilar(i)
	}
	for i := 10000; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}

}
