package main

import (
	"fmt"
	"time"
	us "./user"
)

type pepe struct {
	us.Usuario
}

func main(){
	new_user := new(pepe)
	new_user.InitUsuario(1, "Elías Vasquez", time.Now(), true)
	fmt.Println(new_user.Usuario)
}