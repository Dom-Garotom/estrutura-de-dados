package linkedList

import "fmt"

type LinkedList struct{
	FirstNode *Node
}

/* 
	Função que adiciona um novo nô dentro da nossa lista encadeada.
		Recebe uma string como parâmetro
*/
func (list *LinkedList) AppendNode(data string) {
	var newNode  *Node = &Node{data: data}

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
	Função responsável por apresentar todos os elementos da lista
*/
func (list *LinkedList) ShowAll(){
	var currentNode = list.FirstNode

	if currentNode == nil {
		fmt.Printf("\nA lista está vázia. \n")
		return
	}

	fmt.Printf("\nLista de elementos : \n")

	for currentNode.nextNode != nil {
		fmt.Printf("Elemento - %s\n", currentNode.data)
		currentNode = currentNode.nextNode
	}
}
