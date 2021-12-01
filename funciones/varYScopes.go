package main

import "fmt"

func modifVar(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Cambio de valor"
	fmt.Println("Nuevo valor:", parametro)

	fmt.Printf("%p\n", &parametro)
}

func main() {
	var curso = "Curso de Go!"
	modifVar(curso)
	fmt.Println(">>>", curso)
	fmt.Printf("%p\n", &curso)
}
