package main

import "fmt"

func main() {
	if nombre, edad := "Aky", 96; nombre == "Ako" {
		fmt.Println("Hola", nombre, "Estoy haciendo curso de Go")
	} else {
		fmt.Println("Los valores son: ", nombre, edad)
	}

	//fmt.Println(nombre) undefined
}
