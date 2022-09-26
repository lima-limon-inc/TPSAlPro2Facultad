package lista

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero(T)
	// InsertarUltimo(T)
	// BorrarPrimero() T
	// VerPrimero() T
	// VerUltimo() T
	// Largo() int
	// Iterar(visitar func(T) bool)
	// Iterador() IteradorLista[T]
}
