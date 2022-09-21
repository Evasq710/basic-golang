package main

import "fmt"

/*Interceptores que permiten ejecutar instrucciones comunes
a varias funciones que reciben y devuelven los mismos tipos
de variables*/

var result int

func main(){
	fmt.Println("Inicio")

	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(funcOP func(int, int) int) func(int, int)int{
	return func(a, b int) int {
		// En este espacio se adiciona código a funciones establecidas
		// Aquí se interceptan a las funciones
		fmt.Println("Inicio de operación")
		
		return funcOP(a, b)
	}
}