package main

import (
	"fmt"
)

var tabla [10]int
var matriz [5][7]int

func main(){
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i:= 0; i < len(arr); i++{
		fmt.Println(arr[i])
	}

	matriz[3][1] = 1
	fmt.Println(matriz)

	// Arreglos dinámicos (Slices)
	arr_slice := []int{2, 4, 6}
	fmt.Println(arr_slice)
	
	vector := [5]int{1, 2, 3, 4, 5}
	porcion_slice := vector[2:] // Desde la posición [2] hasta la última, se crea un Slice de un vector
	fmt.Println(porcion_slice)

	slice := make([]int, 5, 20) // make(type, size, capacity) -> make(type, min, espacio reservado)
	fmt.Printf("\nLargo: %d, Capacidad del slice: %d", len(slice), cap(slice))

	slice_vacio := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		// Append recibe un slice y retorna un slice con el elemento agregado
		slice_vacio = append(slice_vacio, i) // usar append cuando supero el largo del slice
	}
	fmt.Printf("\nLargo: %d, Capacidad: %d", len(slice_vacio), cap(slice_vacio)) // La capacidad se restringe a los números binarios (2^x)
}