package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "lista"
	"testing"
)

func TestColaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.EqualValues(t, true, lista.EstaVacia())
}
