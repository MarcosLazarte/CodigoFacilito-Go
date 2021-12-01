package main

import "fmt"

func main() {
	var nota int
	fmt.Print("Ingresa una calificación: ")
	fmt.Scanf("%d", &nota)
	/*
		if nota == 10 {
			fmt.Println("Felicidades! Lograste un 10")
		} else {
			if nota == 8 || nota == 9 {
			} else {
				if nota == 7 {
					fmt.Println("Aprobaste, pero necesitas estudiar un toque más")
				} else {
					if nota >= 0 && nota <= 6 {
						fmt.Println("No aprobaste")
					} else {
						fmt.Println("Calificacion no valida")
					}
				}
			}
		}
	*/
	if nota == 10 {
		fmt.Println("Felicidades! Lograste un 10")
	} else if nota == 8 || nota == 9 {
		fmt.Println("Muy bien, aprobaste")
	} else if nota == 7 {
		fmt.Println("Aprobaste, pero necesitas estudiar un toque más")
	} else if nota >= 0 && nota <= 6 {
		fmt.Println("No aprobaste")
	} else {
		fmt.Println("Calificacion no valida")
	}
}
