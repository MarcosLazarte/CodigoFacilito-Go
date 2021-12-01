package main

import "fmt"

func main() {
	var dividendo, divisor int

	fmt.Print("Ingresa un valor para el dividendo: ")
	fmt.Scanf("%d\n", &dividendo)

	fmt.Print("Ingrese un valor para el divisor: ")
	fmt.Scanf("%d", &divisor)

	if divisor == 0 {
		panic("No es posible dividir sobre 0") //No se ejecuta internamente, lo llamamos nosotros. Validamos un error
	}

	resultado := dividendo / divisor
	fmt.Println("Resultado es: ", resultado)
}
