package main

import "fmt"

func funcion1() {
	fmt.Println("Hola, soy la primera funcion")
}
func funcion2() {
	fmt.Println("Hola, soy la SEGUNDA funcion")
}

func main() {
	fmt.Println("Hola, nos encontramos en la función MAIN!!!")
	defer funcion1()
	funcion2()
}
