package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	// "go" (go rutine) hace que se ejecute de manera asíncrona
	// si la función en la que se llama termina antes, termina el programa
	go nombreLento("Elías")
	
	fmt.Println("Estoy aquí")
	var x string
	fmt.Scanln(&x)
	fmt.Printf("Luego de haber escrito: \"%s\", la ejecución termina.", x)
}

func nombreLento(nombre string){
	letras := strings.Split(nombre, "") // letras se convierte en vector
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond) // pausa de un segundo
	}
}