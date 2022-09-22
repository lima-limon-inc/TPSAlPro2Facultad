package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, true, pila.EstaVacia())
	t.Log("Veo si la pila vacia se sigue comportando como tal incluso despues de apilar y desapilar")
}

func TestPilaPocosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, true, pila.EstaVacia())
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, true, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	t.Log("Hacemos la prueba de volumen")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 10000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, false, pila.EstaVacia())
	}
	for i := 10000; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.EqualValues(t, true, pila.EstaVacia())

	pilaStr := TDAPila.CrearPilaDinamica[string]()
	var alfabeto [26]string = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, letra := range alfabeto {
		pilaStr.Apilar(letra)
	}
	for i := len(alfabeto) - 1; i >= 0; i-- {
		require.EqualValues(t, alfabeto[i], pilaStr.Desapilar())
	}
}
