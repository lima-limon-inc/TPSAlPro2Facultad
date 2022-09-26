package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func nodoCrear[T any](datoArgumento T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = datoArgumento
	nodo.prox = nil
	return nodo
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero = nil
	lista.ultimo = nil
	lista.largo = 0
	return lista
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) insertarListaVacia(elem T) {
	nuevoNodo := nodoCrear(elem)

}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T){
	nuevoNodo := nodoCrear(elem)
	lista.ultimo.prox = nuevoNodo
	ultimo = nuevoNodo
}
