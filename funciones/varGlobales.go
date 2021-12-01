package main

import "fmt"

var user string // inicializa ""

func funcion1() {
	user = "SoyFuncion1"
	fmt.Println("funcion1:", user)
}

func funcion2() {
	user = "SoyFuncion2"
	fmt.Println("funcion2:", user)
}

func main() {
	user = "Aky"
	fmt.Println("main:", user)
	funcion1()
	funcion2()

	fmt.Println("main:", user)
}
