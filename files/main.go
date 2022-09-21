package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
)

func main(){
	leoArchivo()
	leoArchivo2()
	crearArchivo()
	editarArchivo()
}

func leoArchivo(){
	archivo, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2(){
	archivo, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			// lee linea por línea hasta un EOF
			registro := scanner.Text()
			fmt.Printf("Linea > %s\n", registro)
		}
	}
	archivo.Close()
}

func crearArchivo(){
	archivo, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hubo un error")
	} else {
		fmt.Fprintln(archivo, "Esta es una nueva línea")
		archivo.Close()
	}
}

func editarArchivo(){
	if AppendLine("./newFile.txt", "Esta es una segunda línea") {
		fmt.Println("Archivo newFile.txt editado correctamente")
		return
	}
	fmt.Println("Hubo un error en la adición de segunda línea")
}

func AppendLine(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}