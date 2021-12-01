package main

import "fmt"

// Variadic function
func promedio(notas ...int) int {
	var promedio, suma int
	for _, notas := range notas {
		suma = suma + notas
	}
	promedio = suma / len(notas)
	return promedio
}

func main() {
	//Variadic Function
	fmt.Println("Hola", "Mundo", "Desde el curso", "Profesional", "de Go")
	resultado := promedio(10, 4, 7, 7)
	fmt.Println("Promedio:", resultado)
}
