package main

import "fmt"

func main() {
	//numeros := [5]int{100, 200, 300, 400, 500}
	numeros := [...]int{100, 200, 300, 400, 500}
	/*
		var numeros [5]int //Declaro array de numeros enteros
		numeros[0] = 100
		numeros[1] = 200
		numeros[2] = 300
		numeros[3] = 400
		numeros[4] = 500 //Sí asigno "500" dara error

		fmt.Println(numeros[0])
		fmt.Printf("%d\n", numeros[1])
		fmt.Print(numeros[2])
	*/
	fmt.Println(numeros)
}
