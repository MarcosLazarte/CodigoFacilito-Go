package main

import "fmt"

func main() {

	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	/*
		for indice := 0; indice < len(animales); indice++ {

			elemento := animales[indice]
			fmt.Println(elemento)

		}
	*/
	//for indice, elemento =: range animales{
	for _, elemento := range animales { //_ sirve por sí no quiero usar el indice
		//fmt.Println("El indice es: ", indice, "elemento es: ", elemento)
		fmt.Println("El valor es: ", elemento)
	}
}
