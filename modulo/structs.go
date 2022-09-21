package main

import (
	"fmt"
	"time"
)

type usuario struct {
	Id	int
	Nombre	string
	FechaAlta	time.Time
	Status	bool
}

func main(){
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User)
}