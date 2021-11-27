package main

import "fmt"

func main() {
	meses := []string {"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, la capacidad es:  %v %p\n", longitud, capacidad, meses)

	meses = append(meses, "Octubre") //Sí la estructura se encuentra en su capacidad máxima se genera un nuevo array
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitud es: %v, la capacidad es:  %v %p\n", longitud, capacidad, meses)
}
