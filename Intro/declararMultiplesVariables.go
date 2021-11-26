package main

import "fmt"

func main() {
	/*
		var nombre string
		var apellido string
		var pais string
	*/
	//var nombre, apellido, pais string
	//var nombre, apellido, pais = "Aky", "Lazarte", "Argentina"
	nombre, apellido, pais := "Aky", "Lazarte", "Argentina"
	edad, altura := 25, 1.69

	fmt.Println(nombre, apellido, pais, edad, altura)
}
