package main

import "fmt"

func modifValor(parametro *string) {
	*parametro = "Cambio de valor"
}
func main() {
	var curso = "Curso de Go!"
	fmt.Println("Antes de la funcion:", curso)
	modifValor(&curso) //referencia, un espacio en memoria
	fmt.Println("Después de la función:", curso)
}
