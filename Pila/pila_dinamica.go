package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad int //Cantidad de elementos que tiene la pila
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T]) // ACA esta el cambio
	// hago lo que deba hacer
	return pila
}
