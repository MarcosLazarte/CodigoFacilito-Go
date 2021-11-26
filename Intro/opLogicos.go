package main

import "fmt"

func main() {
	/*
		&&
		||
		!
	*/
	//resultado := 20 == 20 && 30 == 30 true
	//resultado := 20 == 20 && 30 == 30 && 20 > 40 false
	resultado := 15 == 15 && 60 < 100 && (90 > 100 || 100 == 90) //false
	falso := !true
	fmt.Println(resultado)
	fmt.Println(falso)
}
