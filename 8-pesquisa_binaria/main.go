package main

import "fmt"

func main() {

	arrayNumbers := [100]int{}
	for i := 0; i < len(arrayNumbers); i++ {
		arrayNumbers[i] = i + 1
	}

	inicio := 0
	fim := len(arrayNumbers) - 1

	continuar := 1

	numeroPesquisa := 47
	comparacoes := 0

	for continuar == 1 {
		comparacoes++
		pivo := (inicio + fim) / 2

		if numeroPesquisa == arrayNumbers[pivo] {
			continuar = 0
			fmt.Println("Elemento encontrado")
		} else if numeroPesquisa > arrayNumbers[pivo] {
			inicio = pivo + 1
		} else {
			fim = pivo - 1
		}

		if fim < inicio {
			continuar = 0
			fmt.Println("Numero não encontrado")
		}

	}

	fmt.Println("A quantidade de comparações foi: ", comparacoes)

}
