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


//Expresion recursiva de la funcion:
//T(n) = 2T(n/2) + O(1)
//log_{2}2 = 1
//C = 0
//C < D
//Ergo ---> O(n^1) = O(n)
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
		return true}
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

//Amongus
func main() {
	a := []int{1, 2, 3,4, 5, 6, 7, 8}
	fmt.Println(ej9(a))
}
