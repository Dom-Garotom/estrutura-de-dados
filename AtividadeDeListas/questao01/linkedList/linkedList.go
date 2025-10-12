/*
Package linkedlist fornece estruturas e funções para manipulação
de listas encadeadas simples.
*/
package linkedlist

import "fmt"

type LinkedList struct {
	FirstNode *Node
}

/*
AppendNode é uma função que adiciona um novo nô dentro da nossa lista encadeada.

	Recebe uma string como parâmetro
*/
func (list *LinkedList) AppendNode(data string) {
	var newNode *Node = &Node{data: data}

	if list.FirstNode == nil {
		list.FirstNode = newNode
		return
	}

	var currentNode *Node = list.FirstNode

	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	currentNode.nextNode = newNode
}

/*
ShowAll é uma função responsável por apresentar todos os elementos da lista
*/
func (list LinkedList) ShowAll() {
	var currentNode = list.FirstNode

	if currentNode == nil {
		fmt.Printf("\nA lista está vázia. \n")
		return
	}

	fmt.Printf("\nLista de elementos : \n")

	for currentNode != nil {
		fmt.Printf("Elemento - %s\n", currentNode.data)
		currentNode = currentNode.nextNode
	}
}

/*
Size é uma função responsável por retornar o tamanho atual da lista.
*/
func (list LinkedList) Size() int {
	var currentNode = list.FirstNode
	var count int

	if currentNode == nil {
		return 0
	}

	for currentNode != nil {
		count++
		currentNode = currentNode.nextNode
	}

	return count
}

/*
FindNode é uma função responsável por procurar um elemento específo dentro de uma lista e retornar esse elemento.
*/
func (list LinkedList) FindNode(findValue string) *Node {
	var currentNode = list.FirstNode

	if findValue == "" || currentNode == nil {
		return nil
	}

	for currentNode != nil {
		if currentNode.data == findValue {
			return currentNode
		}

		currentNode = currentNode.nextNode
	}

	return nil
}

/*
Contains é uma função responsável por informar se um elemento existe em uma lista e retornar um valor booleano correspondente
*/
func (list LinkedList) Contains(node *Node) bool {
	var currentNode = list.FirstNode

	if currentNode == nil || node == nil {
		return false
	}

	for currentNode != nil {
		if currentNode == node {
			return true
		}

		currentNode = currentNode.nextNode
	}

	return false
}

/*
GetValue é uma função responsável por retornar o valor de um nô que tem o acesso privado a outros modulos.
*/
func (list LinkedList) GetValue(node *Node) string {
	return node.data
}

/*
UpdateNode é um função responsável por alterar o valor de um nô presente na lista encadeada.

	Recebe um nô e o novo valor como parâmetros.
*/
func (list LinkedList) UpdateNode(node *Node, value string) {
	var currentNode = list.FirstNode

	if value == "" || currentNode == nil {
		return
	}

	for currentNode != nil {
		if currentNode == node {
			node.data = value
			return
		}

		currentNode = currentNode.nextNode
	}

	fmt.Printf("\nNão foi possível alterar o valor do nô passado\n")
}

/*
GetNextNode é um função responsável por retornar o nô subsequente ao que foi apresentado no parâmetro da função
*/
func (list LinkedList) GetNextNode(node *Node) *Node {
	var currentNode = list.FirstNode

	if node == nil || currentNode == nil {
		return nil
	}

	for currentNode != nil {
		if currentNode == node && currentNode.nextNode != nil {
			return currentNode.nextNode
		}

		currentNode = currentNode.nextNode
	}

	return nil
}

/*
DeleteNode é uma função responsável por excluir um elemento da lista encadeada.

	Recebe como parâmetro o nô que será removido.
*/
func (list *LinkedList) DeleteNode(node *Node) {
	if list.FirstNode == nil || node == nil {
		return
	}

	if list.FirstNode == node {
		list.FirstNode = list.FirstNode.nextNode
		return
	}

	var previousNode *Node = list.FirstNode
	var currentNode *Node = list.FirstNode.nextNode

	for currentNode != nil {
		if currentNode == node {
			previousNode.nextNode = currentNode.nextNode
			return
		}
		previousNode = currentNode
		currentNode = currentNode.nextNode
	}

	fmt.Printf("\nNão foi possível excluir o nô informado, pois ele não existe na lista.\n")
}

/*
Destroy é uma função responsável por destruir toda a lista encadeada, liberando as referências de memória utilizadas.
*/
func (list *LinkedList) Destroy() {
	list.FirstNode = nil
	fmt.Printf("\nLista destruída com sucesso.\n")
}
