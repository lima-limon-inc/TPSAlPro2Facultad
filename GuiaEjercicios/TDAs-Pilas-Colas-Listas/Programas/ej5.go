package main

import "fmt"

type nodoLista[T any] struct {
	prox *nodoLista[T]
	dato T
}
type ListaEnlazada[T any] struct {
	prim *nodoLista[T]
}

func nodoCrear[T any](datoArgumento T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = datoArgumento
	nodo.prox = nil
	return nodo
}

func listaEnlazadaCrear[T any]() ListaEnlazada[T] {
	lista := new(ListaEnlazada[T])
	lista.prim = nil
	return *lista
}

func (lista *ListaEnlazada[T]) Append(n T){
	actual := lista.prim
	if actual == nil {
	lista.prim = nodoCrear(n)

	} else {
	for actual.prox != nil {
		actual = actual.prox
	}
	actual.prox = nodoCrear(n)
}
}

func (lista *ListaEnlazada[T]) Recorrer(k int) T {
	actual := lista.prim
		fmt.Println("K")
	for i := 0; i < k; i++ {
		fmt.Println("D")
		actual = actual.prox
	}
	return actual.dato
}

func main() {
	lista := listaEnlazadaCrear[int]()
	lista.Append(1)
	lista.Append(2)
	lista.Append(3)
	lista.Append(4)
	lista.Append(5)
	lista.Append(6)
	fmt.Println(lista.prim.dato)
	fmt.Println(lista.prim.prox.dato)
	fmt.Println(lista.Recorrer(3))
}
