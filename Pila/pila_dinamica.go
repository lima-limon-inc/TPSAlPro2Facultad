package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad int //Cantidad de elementos que tiene la pila
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	switch pila.cantidad {
	case 0:
		return true
	default:
		return false
	}
}
