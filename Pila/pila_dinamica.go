package pila


/* Definición del struct pila proporcionado por la cátedra. */
const capacidadInicial = 8 //Elijo una potencia de 2 para facilitar las cuentas con enteros a futuro. Sobre todo considerando que voy a multiplicar y dividir por dos

type pilaDinamica[T any] struct {
	datos     []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad  int //Cantidad de elementos que tiene la pila
	capacidad int //Capacidad maxima del array por debajo del array
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, capacidadInicial) //Creo el array con len=0 y cap = capacidadInicial
	pila.cantidad = 0
	pila.capacidad = capacidadInicial
	return pila
}

//func (pila *pilaDinamica[T]) redimensionar(cuanto string) {
//switch cuanto {
//	case "+": //El array tiene que duplicarse
//		pila.capacidad = pila.capacidad * 2
//	case "-": //El array tiene que demediarse
//		pila.capacidad = (pila.capacidad / 2)
//}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	switch pila.cantidad {
	case 0:
		return true
	default:
		return false
	}
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.cantidad == 0 {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1] //Restamos 1 porque la cantidad siempre es mayor, ya que tambien se tiene que contar al item 0 como 1er item
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == pila.capacidad { //Quiero ver si la cantidad va a ser igual a la capacidad de esta funcion asi la redimensiono
		//pila.redimensionar("+")
	}
	pila.datos[pila.cantidad] = elem //Asigno el elemento pasado al ultimo elemento de la lista
	pila.cantidad += 1               //Aumento la cantidad
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.cantidad*4 == pila.capacidad { //Quiero ver si se cumple la condicion de redimension para liberar memoria
		//pila.redimensionar("-")
	}
	// DESAPILAR
	pila.cantidad -= 1 //Aumento la cantidad
	return pila.VerTope()
}
