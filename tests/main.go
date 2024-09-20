package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4", "teste5"}
	evento = evento[:0]

	fmt.Println(evento)
}
