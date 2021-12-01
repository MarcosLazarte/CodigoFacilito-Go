package main

import "fmt"

type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

type Cazador interface {
	Cazar()
}
type Gato struct {
	Name string
}

func (self Gato) Dormir() {
	fmt.Println("El gato duerme")
}
func (self Gato) Comer() {
	fmt.Println("El gato come")
}

func funcionParaUnAnimal(animal Animal) {
	fmt.Println("El objeto es un animal!")
}

func funcionParaUnaMascota(animal Mascota) {
	fmt.Println("El objeto es un mascota!")
}

func funcionparaUnCazador(animal Cazador) {
	fmt.Println("El objeto es un cazador!")
}
func main() {
	gato := Gato{Name: "Mi gato!"}
	gato.Dormir()
	gato.Comer()

	funcionParaUnAnimal(gato)
	funcionParaUnaMascota(gato)
	//funcionparaUnCazador(gato) No tiene el metodo cazar el gato
}
