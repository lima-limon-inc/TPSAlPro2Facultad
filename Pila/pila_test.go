package pila_test

import (
	"fmt"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	fmt.Println(pila.EstaVacia())
}
TestPilaVacia(testing
