package main

import "fmt"

type User struct {
	Name  string // Atributos
	Email string
}

func main() {
	/*
		var aky User // Un objeto
		aky.Name = "Aky"
		aky.Email = "marcos.e.lazarte@gmail.com"
		fmt.Println(aky)
		fmt.Println(aky.Name)
		fmt.Println(aky.Email)
	*/
	//aky := User{Name: "Aky", Email: "marcos.e.lazarte@gmail.com"}
	Name := "Aky"
	Email := "marcos.e.lazarte@gmail.com"
	aky := User{Name, Email}
	fmt.Println(aky.Name)
	fmt.Println(aky.Email)
}
