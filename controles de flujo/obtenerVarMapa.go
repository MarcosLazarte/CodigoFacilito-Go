package main

import "fmt"

func main() {
	usuarios := make(map[string]string)

	usuarios["Aky"] = "marcos.e.lazarte@gmail.com"

	//mail, ok := usuarios["Aky"]

	//if ok {
	if mail, ok := usuarios["Aky"]; ok {
		fmt.Println(mail)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
