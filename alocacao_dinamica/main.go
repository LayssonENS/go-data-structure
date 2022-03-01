package main

import "fmt"

type ponto struct {
	x float32
	y float32
}

func main() {
	var Ponto *ponto

	Ponto = &ponto{
		x: 5,
		y: 8,
	}

	fmt.Println(*Ponto)

}
