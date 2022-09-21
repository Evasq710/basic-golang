package main

import "fmt"

func main() {
	//recibiendo un valor
	fmt.Println(uno(5))
	//recibiendo dos valores
	num, boolean := dos(3)
	fmt.Println(num)
	fmt.Println(boolean)
	//enviando múltiples parámetros
	fmt.Println(Calculo())
	fmt.Println(Calculo(5, 10))
	fmt.Println(Calculo(5, 10, 15, 20))
}

//func nombre (parámetros) lista de tipos de valores a retornar
func uno(numero int) (z int) {
	z = numero*2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// cantidad de parámetros variables
func Calculo(numero ...int) string {
	total, contador := 0, 0
	// range VAR <--- Se usa en elementos iterables, devuelve dos valores: Index, Element
	// _ <------ Aloja los datos que no queremos usar, ahorra uso de memoria
	for _, num := range numero {
		total += num
		contador += 1
	}
	result := fmt.Sprintf("Se recibieron %d números, su suma es %d", contador, total)
	return result
}