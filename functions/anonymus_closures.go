package main

import "fmt"

// Función anónima
var Anonymus_Calc func(int, int) int = func(num1, num2 int) int {
	return num1 + num2
}

func main(){
	fmt.Printf("\nSuma de 5 + 7: %d", Anonymus_Calc(5, 7))

	// Redefiniendo la función anónima
	Anonymus_Calc = func(num1, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("\nResta de 5 - 7: %d\n", Anonymus_Calc(5, 7))

	Operaciones()

	// Closures
	tablaDel := 2
	miTabla := Tabla(tablaDel) // miTabla se convierte en una función
	fmt.Println("==== TABLA DEL 2 ====")
	for i:= 1; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

func Operaciones(){
	resultado := func () int {
		var a int = 3
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// Closures
// Devuelve una función
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0 // variable que solo está en Tabla(), está protegida
	
	return func() int {
		// esta función es la que se ejecutará cada que se llame a miTabla
		secuencia += 1
		return numero * secuencia
	}
}