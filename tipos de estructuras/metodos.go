package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (self *User) setName(name string) { //Metodo
	self.Name = name
}
func (self *User) getName() string {
	return self.Name
}

func main() {
	aky := User{Name: "Aky", Email: "marcos.e.lazarte@gmail.com"}

	aky.setName("Marcos")

	fmt.Println(aky)

	fmt.Println(aky.getName())
}
