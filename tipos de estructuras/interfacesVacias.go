package main

import "fmt"

func mostrarVar(objeto interface{}) {
	fmt.Printf("El valor de la var es: %v\n", objeto)
}
func suma(num1, num2 int) int {
	return num1 + num2
}

type User struct {
	Name string
}

func main() {
	usuario := User{Name: "Aky"}
	mostrarVar(usuario)
}
