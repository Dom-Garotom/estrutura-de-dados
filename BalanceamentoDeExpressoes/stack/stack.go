/*
Package stack é um pacote responsável por gerenciar um pilha de elementos.
Ele oferece funções que ajudam no gerenciamento dessa estrutura de dados.
*/
package stack

import "fmt"

type Stack struct {
	LastIn *Node
}

/*
AddToTop é um função responsável por adicionar um novo item a pilha de elementos.

	Ela recebe uma string como parâmetro.
*/
func (stack *Stack) AddToTop(value string) {
	if value == "" {
		fmt.Printf("\nValores vazios são inválidos")
		return
	}

	currentNode := NewNode(value)
	currentNode.nextNode = stack.LastIn

	stack.LastIn = currentNode
}

/*RemoveFromTop é um função responsável por remover o elemento que fica localizado no topo da pilha de elementos.*/
func (stack *Stack) RemoveFromTop() bool {
	if stack.LastIn == nil {
		fmt.Printf("\nNão foi possível remover o ultimo elemento da pilha porque a pilha  estava vazia")
		return false
	}

	stack.LastIn = stack.LastIn.nextNode
	return true
}
