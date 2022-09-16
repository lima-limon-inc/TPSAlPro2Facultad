package pila_test

import (
	//	"fmt"
	"github.com/stretchr/testify/require"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.EqualValues(t, true, pila.EstaVacia())
}

func TestCantBajaElementos(t *testing.T) {
	t.Log("Hacemos pruebas con una pila de una cantidad baja de elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Apilar(6)
	pila.Apilar(7)
	pila.Apilar(8)
	require.EqualValues(t, 8, pila.VerTope())
	pila.Desapilar()
	pila.Desapilar()
	require.EqualValues(t, 6, pila.VerTope())
}
