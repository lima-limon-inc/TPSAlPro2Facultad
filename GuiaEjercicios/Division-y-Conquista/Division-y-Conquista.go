package main

import "fmt"

func ej10(arreglo[int]) bool {
	if len(arreglo) == 2 {
		if arreglo[0] < arreglo[1] {
			return true
		} else {
			return false
		}
	}
	var mitad int = len(arreglo) / 2
	var izquierda int = mitad - 1
	var derecha int = len(arreglo) - 1
	ej10(arreglo[:izquierda])
	ej10(arreglo[mitad:derecha])
}

func main() {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ej10(a))
}
