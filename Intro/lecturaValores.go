package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu primer nombre: ")
	fmt.Scanf("%s\n", &nombre) //limitado a leer un solo string

	fmt.Print("Ingrese su edad: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Print("Ingrese su altura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con una edad %d y una altura %.2f\n", nombre, edad, altura)
	fmt.Println("Hola", nombre, "con una edad", edad, "y una altura", altura)
}
