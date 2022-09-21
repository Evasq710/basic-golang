package main

import "fmt"

func main(){
	// mapa -> map[clave]valor
	paises := make(map[string]string)
	paises["Guatemala"] = "Guatemala"
	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona": 2,
		"Real Madrid": 4,
		"Boca Juniors": 6}
	fmt.Println(campeonato)

	// agregando elemento
	campeonato["River Plate"] = 25
	// modificando elemento
	campeonato["Barcelona"] = 3
	// eliminando elemento
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, punteo := range campeonato { // RANGE en un MAP devuelve -> clave, valor
		fmt.Printf("El equipo %s, tiene %d puntos\n", equipo, punteo)
	}

	puntaje, existe := campeonato["No existe"] // al llamar a una clave específica, devuelve -> valor, bool si existe
	fmt.Printf("El puntaje capturado es %d, ¿el equipo existe? -> %t \n", puntaje, existe)
	puntaje, existe = campeonato["River Plate"]
	fmt.Printf("El puntaje capturado es %d, ¿el equipo existe? -> %t \n", puntaje, existe)
}