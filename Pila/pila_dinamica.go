package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad int //Cantidad de elementos que tiene la pila
}


