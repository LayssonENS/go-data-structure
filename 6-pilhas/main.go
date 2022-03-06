package main

import (
	"container/list"
	"fmt"
)

type Prato struct {
	Cor string
}

func main() {
	stack := list.New()

	cinza := Prato{"Cinza"}
	azul := Prato{"Azul"}
	verde := Prato{"Verde"}

	stack.PushBack(cinza)
	stack.PushBack(azul)
	stack.PushBack(verde)

	//Imprimir na tela
	Display(stack)

	//Remover elemento da pilha
	stack.Remove(stack.Back())

	fmt.Println("----------------------")

	//Imprimir na tela
	Display(stack)

}
