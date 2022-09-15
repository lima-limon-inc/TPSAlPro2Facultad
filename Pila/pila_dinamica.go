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

func (pila *pilaDinamica[T]) VerTope() int {
	if pila.cantidad == 0 {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad - 1] //Restamos 1 porque la cantidad siempre es mayor, ya que tambien se tiene que contar al item 0 como 1er item
}
