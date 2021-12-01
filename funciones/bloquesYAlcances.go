package main

import "fmt"

func main() { //bloque1

	var curso = "Curso de Go!"
	var version = 1

	{ //bloque2

		fmt.Println(curso)
		//fmt.Println(version)

		{ //bloque3
			var version = 3 //Una var vive dentro de un bloque

			fmt.Println("Version: ", version)
			fmt.Println(curso)

		}
		fmt.Println("Estamos en el bloque2. Version:", version)
	}
	fmt.Println(curso)
}
