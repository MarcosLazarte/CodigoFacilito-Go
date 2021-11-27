package main

import "fmt"

func main() {
	monedas := [...]string {0: "Dólar Canadiense", 3: "Peso Argentino", 2: "Dolar", 5: "Euro"}
	/*
		fmt.Println(monedas[0])
		fmt.Println(monedas[1])
		fmt.Println(monedas[2])
		fmt.Println(monedas[3])
		fmt.Println(monedas[4])
		fmt.Println(monedas[5])
	*/
	subMonedas := monedas[0:3] //(Sub-arreglo) Slice
	fmt.Println(subMonedas)
}
