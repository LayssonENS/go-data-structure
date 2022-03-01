package main

import "fmt"

func main() {
	var ponteiro *int
	valor := 5
	ponteiro = &valor
	fmt.Println("O valor apontado pelo ponteiro é :", *ponteiro)
}
