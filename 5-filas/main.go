package main

import (
	"container/list"
	"fmt"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int8
}

func main() {
	queue := list.New()

	tiago := Pessoa{"Tiago", "Temporin", 31}
	joao := Pessoa{"João", "Santos", 24}
	dani := Pessoa{"Dani", "Almeida", 54}

	queue.PushBack(tiago)
	queue.PushBack(joao)
	queue.PushBack(dani)

	//Imprimir na tela
	Display(queue)

	//Remover elemento da fila
	queue.Remove(queue.Front())

	fmt.Println("----------------------")

	//Imprimir na tela
	Display(queue)

}
