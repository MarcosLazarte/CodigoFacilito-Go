package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
	Cazar()
}
type Perro struct {
	Nombre string
}

func (self Perro) Comer() {
	fmt.Println("El perro come!")
}
func (self Perro) Dormir() {
	fmt.Println("El perro duerme!")
}
func (self Perro) Cazar() {
	fmt.Println("El perro caza!")
}
func ejecutarAcciones(animal Animal) {
	animal.Comer()
	animal.Dormir()
}

func main() {
	ringo := Perro{Nombre: "Ringo"}
	ejecutarAcciones(ringo)
}
