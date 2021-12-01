package main

import "fmt"

func factorial(numero int) int {
	if numero == 1 {
		return 1
	}
	return factorial(numero-1) * numero
}
func main() {
	resultado := factorial(5)
	fmt.Println("Factorial es:", resultado)
}
