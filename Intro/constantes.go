package main

import "fmt"

const Curso string = "Curso profesional de Go!"

func main() {
	//Curso = "Nuevo valor" No se puede re asignar valor de una constante
	fmt.Println(Curso)
}
