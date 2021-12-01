package main

import "fmt"

type Operacion func(balance, cantidad int) int

func suma(numero1, numero2 int) int { //Operacion
	return numero1 + numero2
}
func resta(num1, num2 int) int { //Operacion
	return num1 - num2
}
func multiplicacion(num1, num2 int, num3 int) int { //No funciona por que tiene 3 argumentos
	return num1 * num2 * num3
}
func ejecutarOp(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operación")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es: ", resultado)

	fmt.Println("Desp de la operación")
}

func main() {
	ejecutarOp(resta, 150, 50)
}
