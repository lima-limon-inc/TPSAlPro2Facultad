package pila

/* Definición del struct pila proporcionado por la cátedra. */
const (
	_CAPACIDAD_INICIAL = 8 //Elijo una potencia de 2 para facilitar las cuentas con enteros a futuro. Sobre todo considerando que voy a multiplicar y dividir por dos
	//Siempre dividio pila.capacidad a la mitad. Si quiero dejarla a la mitad, lo multiplico por 1 (redimension); y si quiero el doble, lo multiplico por 4 (redimension). De ahi estos dos valores
	_REDIMENSION_MAYOR = 4 //Esto equivale al doble de tamano
	_REDIMENSION_MENOR = 1 //Esto equivale a la mitad de tamano
)

type pilaDinamica[T any] struct {
	datos     []T //Array de datos de cualquier tipo (todos mismo tipo)
	cantidad  int //Cantidad de elementos que tiene la pila
	capacidad int //Capacidad maxima del array por debajo del array
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.cantidad = 0
	pila.capacidad = _CAPACIDAD_INICIAL
	pila.datos = make([]T, pila.capacidad) //Creo el array con len=0 y cap = _CAPACIDAD_INICIAL
	return pila
}

func (pila *pilaDinamica[T]) redimensionar(cuanto int) {
	nuevoArray := make([]T, pila.capacidad/2*cuanto) //Creo un nuevo array con una distinta capacidad.
	copy(nuevoArray, pila.datos)
	pila.capacidad = cap(nuevoArray)
	pila.datos = nuevoArray

}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0 //Devuelve un true si se cumple y un false si no se cumple
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() == true {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1] //Restamos 1 porque la cantidad siempre es mayor, ya que tambien se tiene que contar al item 0 como 1er item
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == pila.capacidad { //Quiero ver si la cantidad va a ser igual a la capacidad de esta funcion asi la redimensiono
		pila.redimensionar(_REDIMENSION_MAYOR)
	}
	pila.datos[pila.cantidad] = elem //Asigno el elemento pasado al ultimo elemento de la lista
	pila.cantidad += 1               //Aumento la cantidad
}

func (pila *pilaDinamica[T]) Desapilar() T {
	var tope T = pila.VerTope()            //Me guardo la variable antes de nada para paniquear en los casos de pila vacia
	pila.cantidad -= 1                     //Aumento la cantidad
	if pila.cantidad*4 == pila.capacidad { //Quiero ver si se cumple la condicion de redimension para liberar memoria
		pila.redimensionar(_REDIMENSION_MENOR)
	}
	return tope
}
