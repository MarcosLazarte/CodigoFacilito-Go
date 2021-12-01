package main

import "fmt"

func main() {
	var nota int

	fmt.Print("Ingrese una nota: ")
	fmt.Scanf("%d", &nota)
	/*
		switch {
		case nota == 10:
			fmt.Println("Felicidades, lograste una nota perfecta")
		case nota == 9 || nota == 8:
			fmt.Println("Aprobaste")
		case nota == 7:
			fmt.Println("Aprobaste, pero necesitas estudiar más")
		case nota >= 0 && nota <= 6:
			fmt.Println("Desaprobado")
		default:
			fmt.Println("Nota no valida")
		}
	*/
	switch nota {
	case 10:
		fmt.Println("Felicidades, lograste una nota perfecta")
	case 9, 8:
		fmt.Println("Aprobaste")
	case 7:
		fmt.Println("Aprobaste, pero necesitas estudiar más")
	case 0, 1, 2, 3, 4, 5, 6:
		fmt.Println("Desaprobado")
	default:
		fmt.Println("Nota no valida")
	}
}
