package main

import (
	"fmt"
)

func main() {
	values := []int64{3, 6, 8, 20, 9, 2, 6, 8, 9, 3, 5, 40, 7, 9, 13, 6, 8}

	// Inicialize uma árvore de pesquisa binária e adicione elementos
	tree := NewBinarySearchTree()
	for _, v := range values {
		tree.Add(v)
	}

	// Encontre o nó com o valor máximo ou mínimo
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())

	// Encontre 99 que não existe
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}

	// Encontrar 9 existente
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	// Encontrar o 9 e depois excluir
	tree.Delete(9)
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	// Travessia em ordem para alcançar a classificação
	tree.MidOrder()
}
