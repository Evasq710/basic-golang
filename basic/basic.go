package main

// go run basic.go para probar
// go build basic.go para crear el .exe
// nombre de package debe ser el mismo del archivo <------ ?????
// func main es obligatoria, ahí empezará el programa

import (
	"fmt"
	"os"
	"bufio"
)

// 0 por default
var numero int
// "" por default
var texto string
// false por default
var status bool = true

func main(){
	// Otra declaración de variables y casteo
	numero1, numero2, numero3, texto := 1, 2, 3, "Esta es una var local, no el global"
	var casteo int64 = 111
	numero1 = int(casteo)

	fmt.Println(numero)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
	fmt.Println(status)

	// -----------------------------------------

	// mostrar y aceptar datos de dos formas

	fmt.Println("Ingrese número 1:")
	//usando punteros
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese número 2:")
	fmt.Scanln(&numero2)
	fmt.Println("Ingrese descripción:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto = scanner.Text()
	}
	fmt.Println("Números:", numero1, numero2)
	fmt.Println("Descripción:", texto)

	// -----------------------------------------

	// Condicionales

	if status = false; status == true {
		fmt.Println("Status es true")
	}else if status == false {
		fmt.Println("Status es false")
	}

	// switch de una única opción, no necesita break
	switch new_var := 1; new_var {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(new_var, "no está entre 1 y 3")		
	}

	// -----------------------------------------

	// Ciclos
	i := 0
	for i < 10 {
		fmt.Printf("Valor de i: %d \n", i)
		if i == 5 {
			fmt.Printf("Multiplicamos por 2 \n")
			i = i*2
			continue
		}
		fmt.Printf("   Pasó por aquí\n")
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// ciclo infinito
	i = 0
	for {
		fmt.Println("Continúa el ciclo")
		fmt.Println("Ingrese el número secreto (100):")
		fmt.Scanln(&i)
		if i == 100 {
			break
		}
	}

	// goto LABEL
	i = 0
	ETIQUETA:
		fmt.Println("Inicio de label ETIQUETA")
		for i < 10 {
			if i == 4 {
				i += 2
				fmt.Println("goto ETIQUETA sumando 2 a i")
				goto ETIQUETA
			}
			fmt.Printf("Valor de i: %d \n", i)
			i++
		}
}