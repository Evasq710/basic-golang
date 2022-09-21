package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool // al hacer esto, relacionamos "humano" con "serVivo"
}

type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
}

type mujer struct {
	// hereda todos los atributos de hombre
	// al hombre implementar la interfaz humano, mujer también la implementa
	hombre
}

// de esta manera, Hombre implementa la interfaz Humano
func(h *hombre) respirar(){h.respirando = true}
func(h *hombre) comer(){h.comiendo = true}
func(h *hombre) pensar(){h.pensando = true}
func(h *hombre) sexo() string{
	if h.esHombre == true {return "Hombre"}
	return "Mujer"
}
func(h *hombre) estaVivo() bool {return h.respirando}

// recibe cualquier estructura que implemente la interfaz humano
func HumanosRespirando(hum humano){
	hum.respirar()
	hum.pensar()
	fmt.Printf("Soy un(a) %s, y ya estoy respirando\n", hum.sexo())
}

// -------------------------------------

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool // al hacer esto, relacionamos "humano" con "serVivo"
}

type perro struct {
	respirando bool
	comiendo bool
	carnivoro bool
}

func(p *perro) respirar(){p.respirando = true}
func(p *perro) comer(){p.comiendo = true}
func(p *perro) EsCarnivoro() bool{return p.carnivoro}
func(p *perro) estaVivo() bool {return p.respirando}

// recibe cualquier estructura que implemente la interfaz animal
func AnimalesRespirando(anim animal){
	anim.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(anim animal) int{
	if anim.EsCarnivoro() {return 1}
	return 0
}

// ---------------------------

// recibe cualquier estructura que implemente la interfaz Ser Vivo
// (humanos y animales)
func EstoyVivo(ser serVivo) bool{
	return ser.estaVivo()
}


func main(){
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)
	fmt.Printf("Estoy vivo: %t\n", EstoyVivo(Pedro))

	Maria := new(mujer)
	HumanosRespirando(Maria)
	fmt.Printf("Estoy vivo: %t\n", EstoyVivo(Maria))

	//---------------------
	totalCarnivoros := 0
	
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirando(Dogo)
	fmt.Printf("Estoy vivo: %t\n", EstoyVivo(Dogo))
	totalCarnivoros += AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnívoros: %d\n", totalCarnivoros)
}