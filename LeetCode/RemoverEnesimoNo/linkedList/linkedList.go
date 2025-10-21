/*
Package linkedlist fornece estruturas e funções para manipulação
de listas encadeadas simples.
*/
package linkedlist

import "fmt"

type LinkedList struct {
	FirstNode *Node
}

func RemoveNthNodeFromEndList(head *LinkedList, removeSteps int) *LinkedList {
	removePointer := head.FirstNode
	advancerPointer := head.FirstNode

	// adianta o advancerPointer a quantidade de passos passados do removeSteps
	for range removeSteps {
		if advancerPointer == nil {
			return &LinkedList{head.FirstNode.NextNode}
		}

		advancerPointer = advancerPointer.NextNode
	}

	// percorre a lista com o removePointer N possições atrás do advancedPointer para que ele para no ponto correto de remoção
	for advancerPointer != nil {
		if advancerPointer.NextNode != nil {
			removePointer.NextNode = removePointer.NextNode.NextNode
			return head
		}

		removePointer = removePointer.NextNode
		advancerPointer = advancerPointer.NextNode
	}

	head.FirstNode = head.FirstNode.NextNode
	return head
}

func (list *LinkedList) Size() int {
	currentNode := list.FirstNode

	if currentNode == nil {
		return 0
	}

	var count int

	for currentNode != nil {
		count++
		currentNode = currentNode.NextNode
	}

	return count
}

func (list *LinkedList) AppendNode(data int) {
	newNode := &Node{Value: data}

	if list.FirstNode == nil {
		list.FirstNode = newNode
		return
	}

	currentNode := list.FirstNode

	for currentNode.NextNode != nil {
		currentNode = currentNode.NextNode
	}

	currentNode.NextNode = newNode
}

func (list LinkedList) ShowAll() {
	var currentNode = list.FirstNode

	if currentNode == nil {
		fmt.Printf("[]\n")
		return
	}

	fmt.Printf("[")

	for currentNode != nil {
		if currentNode.NextNode == nil {
			fmt.Printf("%d", currentNode.Value)
			break
		}

		fmt.Printf("%d, ", currentNode.Value)
		currentNode = currentNode.NextNode
	}

	fmt.Printf("]\n")
}
