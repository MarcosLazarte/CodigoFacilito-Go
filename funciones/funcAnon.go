package main

import "fmt"

func main() {
	/*
		func() {
			fmt.Println("Hola desde func sin nombre!")
		}()
	*/
	/*
		miFuncion := func(user string) {
			fmt.Println("Hola", user, " desde func sin nombre!")
		}
		miFuncion("Aky")
		miFuncion("Aky2")
		miFuncion("Aky3")
	*/
	miFuncion := func(user string) string {
		message := fmt.Sprintf("Hola %s, te saludamos desde una función sin nombre", user)
		return message
	}
	mensajeFinal := miFuncion("Aky")
	fmt.Println(mensajeFinal)
}
