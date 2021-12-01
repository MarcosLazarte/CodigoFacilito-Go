package main

import "fmt"

func main() {
	usuarios := map[int]string{} // make
	usuarios[0] = "Usuario 1"
	usuarios[1] = "Usuario 2"
	usuarios[2] = "Usuario 3"
	usuarios[3] = "Usuario 4"

	for id, usuario := range usuarios {
		fmt.Println(id, usuario)
	}
}
