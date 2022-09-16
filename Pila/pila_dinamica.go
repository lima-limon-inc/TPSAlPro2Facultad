package pila

import "fmt"

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
	//fmt.Println(pila)
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
	fmt.Println("REDIMENSIONAR")
	nuevoArray := make([]T, pila.capacidad/2*redimension) //Creo un nuevo array con una distinta capacidad. Siempre dividio pila.capacidad a la mitad. Si quiero dejarla a la mitad, lo multiplico por 1 (redimension); y si quiero el doble, lo multiplico por 4 (redimension)
	fmt.Println("ANTES")
	fmt.Println("nuevoArray")
	fmt.Println(nuevoArray)
	fmt.Println("pila.datos")
	fmt.Println(pila.datos)
	copy(nuevoArray, pila.datos)
	fmt.Println("DESPUES")
	fmt.Println("nuevoArray")
	fmt.Println(nuevoArray)
	fmt.Println("pila.datos")
	fmt.Println(pila.datos)
	//fmt.Println("CAMBIO EL ARRAY DE LA PILA")
	//fmt.Println("Antes")
	//fmt.Println(pila)
	//fmt.Println("Despues")
	pila.capacidad = cap(nuevoArray)
	pila.datos = nuevoArray
	//fmt.Println(pila)

}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	//fmt.Println("LLAMO A ESTAVACIA")
	//fmt.Println(pila)
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
	//fmt.Println("LLAMO A APILAR")
	if pila.cantidad == pila.capacidad { //Quiero ver si la cantidad va a ser igual a la capacidad de esta funcion asi la redimensiono
		//fmt.Println("redimensionar")
		pila.redimensionar("+")
	}
	pila.datos[pila.cantidad] = elem //Asigno el elemento pasado al ultimo elemento de la lista
	pila.cantidad += 1               //Aumento la cantidad
	fmt.Println(pila)
}

func (pila *pilaDinamica[T]) Desapilar() T {
	//fmt.Println("LLAMO A DESAPILAR")
	if pila.cantidad*4 == pila.capacidad { //Quiero ver si se cumple la condicion de redimension para liberar memoria
		//fmt.Println("redimensionar")
		pila.redimensionar("-")
	}
	// DESAPILAR
	pila.cantidad -= 1 //Aumento la cantidad
	fmt.Println(pila)
	return pila.VerTope()
}
