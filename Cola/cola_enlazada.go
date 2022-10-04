package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func nodoCrear[T any](datoArgumento T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = datoArgumento
	nodo.prox = nil
	return nodo
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")

	}
	return cola.primero.dato

}
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodoNuevo := nodoCrear(elem) //Creo un nodoNuevo con el dato pasado
	if cola.EstaVacia() {
		cola.primero = nodoNuevo
	} else {
		cola.ultimo.prox = nodoNuevo //Hago que el ultimo elemento de la lista apunte a ese nuevo elemento
	}
	cola.ultimo = nodoNuevo //Hago que el puntero ultimo apunte a la direccion de nodoNuevo. nodoNuevo es un puntero a un nodo

}

func (cola *colaEnlazada[T]) Desencolar() T {
	frenteCola := cola.VerPrimero() // Este caso casa los casos de que se quiera desencolar una cola vacia
	cola.primero = cola.primero.prox
	if cola.primero == nil { //En caso que desencolar vacie la lista
		cola.ultimo = nil
	}
	return frenteCola

}
