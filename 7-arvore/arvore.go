package main

import "fmt"

// BinarySearchTree nó da árvore de pesquisa binária
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // 树根节点
}

// BinarySearchTreeNode nó da árvore de busca binária
type BinarySearchTreeNode struct {
	Value int64                 // valor
	Times int64                 // o número de vezes que o valor aparece
	Left  *BinarySearchTreeNode // subárvore esquerda
	Right *BinarySearchTreeNode // Árvore de palavras à direita
}

// NewBinarySearchTree Inicializa uma árvore de busca binária
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// Add adiciona elemento
func (tree *BinarySearchTree) Add(value int64) {
	//Se não houver raiz, prova ser uma árvore vazia, adiciona a raiz e retorna
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}

	// adiciona o valor em
	tree.Root.Add(value)
}

func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		// Se o valor inserido for menor que o valor do nó, insira-o na subárvore esquerda do nó
		// Se a subárvore esquerda estiver vazia, adicione-a diretamente
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			// senão recua
			node.Left.Add(value)
		}
	} else if value > node.Value {
		// Se o valor inserido for maior que o valor do nó, insira-o na subárvore direita do nó
		// Se a subárvore direita estiver vazia, adicione-a diretamente
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			// senão recua
			node.Right.Add(value)
		}
	} else {
		// O valor é o mesmo, não precisa adicionar, basta adicionar 1 ao número de ocorrências do valor
		node.Times = node.Times + 1
	}
}

// FindMinValue encontra o nó com o menor valor
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		// Se for uma árvore vazia, retorne vazia
		return nil
	}

	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	// A subárvore esquerda está vazia, a superfície já é o nó mais à esquerda e este valor é o valor mínimo
	if node.Left == nil {
		return node
	}

	// Sempre recurse na subárvore esquerda
	return node.Left.FindMinValue()
}

// FindMaxValue Encontre o nó com o valor máximo
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		// Se for uma árvore vazia, retorne vazia
		return nil
	}

	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	// A subárvore direita está vazia, a superfície já é o nó mais à direita e este valor é o valor máximo
	if node.Right == nil {
		return node
	}

	// Sempre recurse na subárvore direita
	return node.Right.FindMaxValue()
}

// Find Encontre o nó especificado
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// Se for uma árvore vazia, retorne vazia
		return nil
	}

	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if value == node.Value {
		// Se o nó for igual ao valor, então retorne o nó
		return node
	} else if value < node.Value {
		// Se o valor pesquisado for menor que o valor do nó, inicie a pesquisa na subárvore esquerda do nó
		if node.Left == nil {
			//A subárvore esquerda está vazia, indicando que o valor não pode ser encontrado e retorna nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// Se o valor pesquisado for maior que o valor do nó, inicie a pesquisa na subárvore direita do nó
		if node.Right == nil {
			// A subárvore direita está vazia, indicando que o valor não pode ser encontrado e retorna nil
			return nil
		}
		return node.Right.Find(value)
	}
}

// FindParent Encontre o pai do nó especificado
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	// Se o nó raiz for igual a esse valor, o nó raiz não tem nó pai e retornará nil
	if tree.Root.Value == value {
		return nil
	}
	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	// Não há determinação de igualdade de valor na camada externa, porque o nó
	// pai é retornado após a determinação da camada interna.

	if value < node.Value {
		// Se o valor pesquisado for menor que o valor do nó, inicie a pesquisa na subárvore esquerda do nó
		leftTree := node.Left
		if leftTree == nil {
			// A subárvore esquerda está vazia, indicando que o valor não pode ser encontrado e retorna nil
			return nil
		}

		// O valor do nó raiz da subárvore esquerda é exatamente igual a esse valor, então o pai é o nó atual, retorne
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}
	} else {
		// Se o valor pesquisado for maior que o valor do nó, inicie a pesquisa na subárvore direita do nó
		rightTree := node.Right
		if rightTree == nil {
			// A subárvore direita está vazia, indicando que o valor não pode ser encontrado e retorna nil
			return nil
		}

		// O valor do nó raiz da subárvore direita é exatamente igual a esse valor, então o pai é o nó atual, retorne
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

// Delete . excluir o elemento especificado
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		// Se for uma árvore vazia, retorne diretamente
		return
	}

	// Descubra se o valor existe
	node := tree.Root.Find(value)
	if node == nil {
		// Se o valor não existir, retorne diretamente
		return
	}

	// Encontre o nó pai do valor
	parent := tree.Root.FindParent(value)

	// No primeiro caso, o nó raiz é excluído e o nó raiz não tem filhos
	if parent == nil && node.Left == nil && node.Right == nil {
		// Retorne diretamente após o esvaziamento
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		//  No segundo caso, o nó excluído tem um nó pai, mas nenhuma subárvore

		// Se o nó a ser deletado for o filho esquerdo do pai, basta deletar o valor diretamente
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			// O original deletado é filho direito do pai, basta deletar o valor diretamente
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// No terceiro caso, existem duas subárvores sob o nó excluído, pois o valor da subárvore direita é maior que
		// o da subárvore esquerda, então o nó excluído é substituído pelo menor elemento da subárvore direita.
		// natureza da árvore de busca binária Satisfeito novamente.

		//  Encontre o menor valor na subárvore direita e continue indo para a esquerda da subárvore direita
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// excluir o menor nó
		tree.Delete(minNode.Value)

		// O nó com o menor valor substitui o nó excluído
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// No quarto caso, há apenas uma subárvore, então a subárvore pode substituir diretamente o nó excluído

		// Se o pai estiver vazio, significa que o nó raiz foi excluído e a raiz da árvore foi substituída.
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// A subárvore esquerda não está vazia
		if node.Left != nil {
			// Se o nó excluído for o filho esquerdo do pai, deixe a subárvore esquerda do nó excluído assumir
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// Se o nó excluído for o filho esquerdo do pai, deixe a subárvore direita do nó excluído assumir
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// MidOrder . Percurso em ordem
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}

	// Imprima a subárvore esquerda primeiro
	node.Left.MidOrder()

	// Imprima o nó raiz por número de vezes
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// imprimir subárvore direita
	node.Right.MidOrder()
}
