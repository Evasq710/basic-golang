package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	// ejemploDefer()
	ejemploPanic()
}

func ejemploDefer(){
	file, err := os.Open("prueba.txt")

	// Se ejecutará cuando se salga de la función
	defer file.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}
}

func ejemploPanic(){
	defer func() {
		reco := recover() // recover trae el resultado de un panic
		// si no hubo panic, reco = nil
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic\n %v", reco)
		}
	}()

	if a := 1; a == 1 {
		panic("Panic: Se encontró el valor de 1")
	}
}