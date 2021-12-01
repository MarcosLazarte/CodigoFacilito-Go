package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return factorial(num-1) * num
}

type customFunc func(n int) int

func main() {
	//var miFunc = factorial
	//var miFunc func(n int) int
	var miFunc customFunc
	if miFunc == nil {
		miFunc = factorial
	}

	resultado := miFunc(3)

	fmt.Println(resultado)
}
