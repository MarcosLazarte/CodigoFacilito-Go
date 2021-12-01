package main

import "fmt"

//func suma(numero1 int, numero2 int) int {
//func suma(numero1, numero2 int) int {
func suma(numero1, numero2 int) (int, string) {
	return numero1 + numero2, "Un mensj desde la func suma"
}

func main() {
	resultado, _ := suma(1, 2)
	fmt.Println("El resultado es: ", resultado)
}
