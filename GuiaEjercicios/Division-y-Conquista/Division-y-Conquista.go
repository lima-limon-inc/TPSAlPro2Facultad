package main

import "fmt"

func ej10(arreglo []int) int {
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	var mitad int = len(arreglo) / 2
	var izquierda int = mitad - 1
	var derecha int = len(arreglo) - 1
	valorI := ej10(arreglo[:izquierda])
	valorD := ej10(arreglo[mitad:derecha])
	if valorI > valorD {
		return 1
	} else {
		return 0
	}
}

// Expresion recursiva de la funcion:
// T(n) = 2T(n/2) + O(1)
// log_{2}2 = 1
// C = 0
// C < D
// Ergo ---> O(n^1) = O(n)
func ej8(arreglo []int) int {
	var min int = arreglo[0]
	if len(arreglo) == 1 {
		if arreglo[0] < min {
			return arreglo[0]
		} else {
			return min
		}
	}
	var mitad int = len(arreglo) / 2
	var izquierda int = mitad
	var derecha int = len(arreglo)
	valorI := ej8(arreglo[:izquierda])
	//	fmt.Println(valorI)
	valorD := ej8(arreglo[mitad:derecha])
	//	fmt.Println(valorD)
	if valorD > valorI {
		return valorI
	} else {
		return valorD
	}

}

func ej9(arreglo []int, principio int, final int) bool {
	if principio == final {
		return true
	}
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	var mitad int = len(arreglo) / 2
	var izquierda int = mitad
	var derecha int = len(arreglo)
	valorI := ej9(arreglo[:izquierda])
	//		fmt.Println(valorI)
	valorD := ej9(arreglo[mitad:derecha])
	//		fmt.Println(valorD)
	if valorI && valorD == true {
		return true
	} else {
		return false
	}
}



func ej11(arreglo []int, extra int) int {
	medio := len(arreglo/2)
	izquierda := arreglo[0:medio]
	derecha := arreglo[medio:len(arreglo)]
	var mitad int = len(izquierda) - 1
	if mitad ==0{
		return mitad
} else if derecha[0] == 1{
		ej11(derecha, mitad)
	} else if izquierda[0] == 1 && izquierda[mitad]{
		return extra + mitad
	}


}

// if f(b) > f(a) --> f(x) = - f(x) Invierto la funcion para que se comporte igual que si f(a) > f(b)
// if f(a) > f(b) or f(b) > f(a)
// div y con
// div y con:
// c = (b + a) / 2
// if f(c) == 0 --> return c
// if f(c) > 0 --> div y con [c:b]
// if f(c) < 0 --> div y con [a:c]
func raiz(f(int) int, a int, b int) int {
	var imagenA int = f(a)
	var imagenB int = f(b)
	if (imagenA > 0 && imagenB < 0) || (imagenA < 0 && imagenB > 0) {


	}
}

func polinomio (a int){
	var rta int = 3 * a
	return rta

}

// Amongus
func main() {
	raiz(polinomio(4),4,5)
}
