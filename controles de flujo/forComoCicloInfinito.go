package main

import "fmt"

func main() {
	var numero = 1
	for {
		fmt.Println(numero) //Se sale con control+c
		numero++

		if numero == 100 {
			//break
			panic("Fin del ciclo for")
		}
	}
}
