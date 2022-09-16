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
	pila.cantidad = 0
	pila.capacidad = capacidadInicial
	pila.datos = make([]T, pila.capacidad) //Creo el array con len=0 y cap = capacidadInicial
	return pila
}

func (pila *pilaDinamica[T]) redimensionar(cuanto string) {
	var redimension int //Variable que me va indicar cuanto redimensionar
	switch cuanto {
	case "+": //El array tiene que duplicarse
		redimension = 4
	case "-": //El array tiene que demediarse
		redimension = 1
	}
	nuevoArray := make([]T, pila.capacidad/2*redimension) //Creo un nuevo array con una distinta capacidad. Siempre dividio pila.capacidad a la mitad. Si quiero dejarla a la mitad, lo multiplico por 1 (redimension); y si quiero el doble, lo multiplico por 4 (redimension)
	copy(nuevoArray, pila.datos)
	pila.capacidad = cap(nuevoArray)
	pila.datos = nuevoArray

}

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
		pila.redimensionar("+")
	}
	pila.datos[pila.cantidad] = elem //Asigno el elemento pasado al ultimo elemento de la lista
	pila.cantidad += 1               //Aumento la cantidad
}

func (pila *pilaDinamica[T]) Desapilar() T {
	var tope T = pila.VerTope()            //Me guardo la variable antes de nada para paniquear en los casos de pila vacia
	if pila.cantidad*4 == pila.capacidad { //Quiero ver si se cumple la condicion de redimension para liberar memoria
		pila.redimensionar("-")
	}
	pila.cantidad -= 1 //Aumento la cantidad
	return tope
}
