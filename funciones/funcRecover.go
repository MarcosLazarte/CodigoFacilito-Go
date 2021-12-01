package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("El programa se cerro de manera incorrecta") //Ya no sale con status 2
		}
	}()

	if file, err := os.Open("examples.txt"); err != nil {
		panic("No fue posible obtener el archivo")
	} else {
		defer func() {
			fmt.Println("\nCerramos el archivo!")
			file.Close()
		}()
		contenido := make([]byte, 254)

		longitud, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo)
	}

}
