package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}
type Estudiante struct {
	User User
	Id   string
}

func main() {
	marcos := User{Name: "Marcos", Email: "marcos.e.lazarte@gmail.com", Active: true}

	melina := User{Name: "Melina", Email: "MelB@gmail.com", Active: true}

	marcosEstudiante := Estudiante{User: marcos, Id: "1f1"}

	fmt.Println(melina)
	fmt.Println(marcosEstudiante.User.Name)
}
