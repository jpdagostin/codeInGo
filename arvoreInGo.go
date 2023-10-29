package main

import "fmt"

// Definindo uma estrutura de um Nó na árvore
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Função para criar um novo nó
func createNewNode(valor int) *Node {
	newNode := &Node{Value: valor, Left: nil, Right: nil}
	return newNode
}

// Função para inserir um novo nó na árvore
func insertInBST(raiz *Node, valor int) *Node {
	if raiz == nil {
		return createNewNode(valor)
	}
	if valor < raiz.Value {
		raiz.Left = insertInBST(raiz.Left, valor)
	} else if valor > raiz.Value {
		raiz.Right = insertInBST(raiz.Right, valor)
	}
	return raiz
}

//Função para percorrer a árvore EM_ORDEM
func inOrder(raiz *Node) {
	if raiz != nil {
		inOrder(raiz.Left)
		fmt.Println(raiz.Value)
		inOrder(raiz.Right)
	}
}

//Função para percorrer a árvore Pré_ORDEM
func preOrder(raiz *Node) {
	if raiz != nil {
		fmt.Println(raiz.Value)
		preOrder(raiz.Left)
		preOrder(raiz.Right)
	}
}

//Função para percorrer a árvore Pós_ORDEM
func posOrder(raiz *Node) {
	if raiz != nil {
		posOrder(raiz.Left)
		posOrder(raiz.Right)
		fmt.Println(raiz.Value)
	}
}

//Função para buscar um Nó
func searchOneNode(raiz *Node, valor int)*Node{
	if raiz == nil{
		fmt.Println("Não foram encontrados valores")
	}
	if raiz.Value == valor{
		fmt.Println("Valor encontrado:" raiz.Value)
		return raiz
	}
	if valor < raiz.Value {
		raiz.Left = searchOneNode(raiz.Left, valor)
	} else if valor > raiz.Value {
		raiz.Right = searchOneNode(raiz.Right, valor)
	}
	return raiz
}

// Função para excluir um nó folha
func deleteSheetNode(raizOld, raiz *Node, valor int) *Node {
    if raiz == nil {
        return raiz
    }
    
    if valor < raiz.Value {
        raiz.Left = deleteSheetNode(raiz, raiz.Left, valor)
    } else if valor > raiz.Value {
        raiz.Right = deleteSheetNode(raiz, raiz.Right, valor)
    } else {
        // Valor encontrado, verificar se é um nó folha
        if raiz.Left == nil && raiz.Right == nil {
            if raizOld != nil {
                if raizOld.Left == raiz {
                    raizOld.Left = nil
                } else if raizOld.Right == raiz {
                    raizOld.Right = nil
                }
            }
            return nil
        }
		else{
			fmt.Println("Este valor não é um nó folha")
		}
    }
    return raiz
}

// Função para excluir um nó com 1 filho
func deleteNodeWithOneSon(raizOld, raiz *Node, valor int) *Node {
    if raiz == nil {
        return raiz
    }

    if valor < raiz.Value {
        raiz.Left = deleteNodeWithOneSon(raiz, raiz.Left, valor)
    } else if valor > raiz.Value {
        raiz.Right = deleteNodeWithOneSon(raiz, raiz.Right, valor)
    } else {
        // Valor encontrado, verificar qual nó (esquerda ou direita) será excluído
        if raiz.Left != nil && raiz.Right == nil {
            // Se o nó a ser excluído tem apenas um filho à esquerda
            temp := raiz.Left
            return temp
        } else if raiz.Left == nil && raiz.Right != nil {
            // Se o nó a ser excluído tem apenas um filho à direita
            temp := raiz.Right
            return temp
        } else {
            // O nó a ser excluído não tem filhos
            if raizOld != nil {
                if raizOld.Left == raiz {
                    raizOld.Left = nil
                } else if raizOld.Right == raiz {
                    raizOld.Right = nil
                }
            }
            return nil
        }
    }
    return raiz
}

// Função para encontrar o menor valor na árvore
func findMinValue(node *Node) *Node {
    current := node
    for current != nil && current.Left != nil {
        current = current.Left
    }
    return current
}


// Função para excluir um nó com dois filhos
func deleteNodeWithTwoSon(raizOld, raiz *Node, valor int) *Node {
    if raiz == nil {
        return raiz
    }

    if valor < raiz.Value {
        raiz.Left = deleteNodeWithTwoSon(raiz, raiz.Left, valor)
    } else if valor > raiz.Value {
        raiz.Right = deleteNodeWithTwoSon(raiz, raiz.Right, valor)
    } else {
        // Valor encontrado na raiz atual.
        if raiz.Left == nil {
            // Se a raiz atual não tem um filho à esquerda, atribua o filho à direita (temp) à raiz atual.
            temp := raiz.Right
            return temp
        }
        if raiz.Right == nil {
            // Se a raiz atual não tem um filho à direita, atribua o filho à esquerda (temp) à raiz atual.
            temp := raiz.Left
            return temp
        }

        // O nó a ser excluído tem dois filhos.
        temp := findMinValue(raiz.Right)

        // Substitua o valor da raiz pelo valor do nó mínimo da subárvore à direita.
        raiz.Value = temp.Value

        // Recursivamente, exclua o nó mínimo da subárvore à direita.
        raiz.Right = deleteNodeWithTwoSon(raiz, raiz.Right, temp.Value)
    }
    return raiz
}

func main() {
    var raiz *Node
    option := 0
    value := 70
    removeValue1 := 80
    removeValue2 := 70
    removeValue3 := 60

    raiz = insertInBST(raiz, 50)
    raiz = insertInBST(raiz, 30)
    raiz = insertInBST(raiz, 20)
    raiz = insertInBST(raiz, 40)
    raiz = insertInBST(raiz, 70)
    raiz = insertInBST(raiz, 60)
    raiz = insertInBST(raiz, 80)
    raiz = insertInBST(raiz, 69)

    for {
        fmt.Println("\n----------------------------------")
        fmt.Println("\nMenu:\n\n")
        fmt.Println("1. Excluir um nó folha\n")
        fmt.Println("2. Excluir nó com uma folha\n")
        fmt.Println("3. Excluir nó com duas folhas\n")
        fmt.Println("4. Imprimir Árvore\n")
        fmt.Println("5. Sair\n")
        fmt.Println("\n----------------------------------\n")

        fmt.Println("Escolha a opção desejada:")
        fmt.Scan(&option)

        switch option {
        case 1:
            fmt.Println("Encontrei o valor a ser removido: ", deleteSheetNode(raiz, raiz, removeValue1))
        case 2:
            fmt.Println("Encontrei o valor a ser removido: ", deleteNodeWithOneSon(raiz, raiz, removeValue3))
        case 3:
            fmt.Println("Encontrei o valor a ser removido: ", deleteNodeWithTwoSon(raiz, raiz, removeValue2))
        case 4:
            fmt.Println("Árvore em in-ordem:")
            inOrder(raiz)
            fmt.Println("\n\n")
            fmt.Println("Árvore em pre-ordem:")
            preOrder(raiz)
            fmt.Println("\n\n")
            fmt.Println("Árvore em pos-ordem:")
            posOrder(raiz)
        case 5:
            fmt.Println("Saindo do Programa")
            return
        default:
            fmt.Println("Opção inválida, tente novamente ...")
        }
    }
}
