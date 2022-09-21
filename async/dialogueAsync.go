package main

import (
	"fmt"
	"time"
)

func main(){
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegué hasta aquí")

	// "AWAIT", hasta que no se cumpla esta promesa, no continúa
	msg := <- canal1
	fmt.Println(msg)

	fmt.Println("Final de ejecución, el canal devolvió respuesta")
}

func bucle(canal1 chan time.Duration){
	inicio := time.Now()

	for i := 0; i < 10000000000; i++ {}

	final := time.Now()

	canal1 <- final.Sub(inicio) // devuelve la duración de la ejecución
}