/*
Package stack é um pacote responsável por gerenciar um pilha de elementos.
Ele oferece funções que ajudam no gerenciamento dessa estrutura de dados.
*/
package stack

import "fmt"

type Stack struct {
	lastIn *Node
}

func CreateStack() *Stack {
	return &Stack{}
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
	currentNode.nextNode = stack.lastIn

	stack.lastIn = currentNode
}

/*RemoveFromTop é um função responsável por remover o elemento que fica localizado no topo da pilha de elementos.*/
func (stack *Stack) RemoveFromTop() bool {
	if stack.lastIn == nil {
		fmt.Printf("\nNão foi possível remover o ultimo elemento da pilha porque a pilha  estava vazia")
		return false
	}

	stack.lastIn = stack.lastIn.nextNode
	return true
}

/*
ViewTheTop é um função responsável por mostrar o valor do elemento localizado no topo da pilha de elementos.

	Retorna uma string contendo o valor do elemento presente no topo da pilha
*/
func (stack *Stack) ViewTheTop() string {
	if stack.lastIn == nil {
		fmt.Printf("\nA lista está vazia")
		return ""
	}

	return stack.lastIn.data
}

/*
IsEmpty é uma função responsável por informar se a pilha de elemento está vazia.

	retorna um bool com resultado da verificação.
*/
func (stack *Stack) IsEmpty() bool {
	return stack.lastIn == nil
}
