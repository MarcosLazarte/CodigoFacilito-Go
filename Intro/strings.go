package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var curso string = "Curso de Go"
	//var curso = "Curso de Go"
	curso := "Curso de Go"
	longitud := len(curso) //int
	fmt.Println(longitud)

	primerCaracter := curso[0] //char

	fmt.Println(primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter))
	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c", curso[0])
}
