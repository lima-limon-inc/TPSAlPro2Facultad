package cola_test

import (
	TDACola "cola"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con una cola de una cantidad baja de elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	require.EqualValues(t, true, cola.EstaVacia())
	cola.Encolar(1)
	cola.Encolar(3)
	require.EqualValues(t, false, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	// require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}
