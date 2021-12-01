package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	User string
	Type UserType
}

func main() {
	aky := User{User: "Aky", Type: Student}
	meli := User{User: "Melina", Type: Teacher}

	fmt.Println(aky)
	fmt.Println(meli)

	if aky.Type == Student {
		fmt.Println("El usuario", aky.User, "es un estudiante")
	}
}
