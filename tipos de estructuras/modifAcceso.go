package main

import (
	"fmt"

	"./CodigoAko"
)

func main() {
	curso := CodigoAky.Curso{Titulo: "Curso Go!"}

	fmt.Println(curso.ObtenerTitulo())
}
