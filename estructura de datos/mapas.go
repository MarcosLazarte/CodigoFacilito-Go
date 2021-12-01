package main

import "fmt"

func main() {
	/*
		dias := make(map[int]string)

		dias[0] = "Domingo"
		dias[1] = "Lunes"
		dias[2] = "Martes"
		dias[3] = "Miercoles"
		dias[4] = "Jueves"
		dias[5] = "Viernes"
		dias[6] = "Sabado"

		dias[4] = "JUEVES"
		delete(dias, 4)
		fmt.Println(dias)
	*/
	usuarios := make(map[string][]int)
	usuarios["Aky"] = []int{10, 9, 8, 10, 5}
	usuarios["Mermelada"] = []int{10, 10, 10, 10, 10}

	fmt.Println(usuarios)
}
