package pila

/* Definición del struct pila proporcionado por la cátedra. */
const capacidadInicial = 16 //Elijo una potencia de 2 para facilitar las cuentas con enteros a futuro. Sobre todo considerando que voy a duplicar (x2) y dividir por 4 (/ 2 x 2)

type pilaDinamica[T any] struct {
	datos    []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad int //Cantidad de elementos que tiene la pila
	capacidad int //Capacidad maxima del array por debajo del array
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	return pila
}

func (pila *pilaDinamica[T]) Apilar() {


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

func (pila *pilaDinamica[T]) Apilar() {
	if pila.cantidad == pila.capacidad { //Quiero ver si la cantidad va a ser igual a la capacidad de esta funcion asi la redimensiono
		pila.redimensionar("+")
	} 
}
