package main

import "fmt"

type Op func(balance, cantidad int) int

func crearOp(tipo string) Op {
	if tipo == "suma" {
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo == "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}
}

func main() {
	suma := crearOp("multi")

	resultado := suma(40, 50)

	fmt.Println("Resultado es: ", resultado)
}
