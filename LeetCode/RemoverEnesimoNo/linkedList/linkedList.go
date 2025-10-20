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
	currentNode := head.FirstNode
	listSize := head.Size()
	removePoint := listSize - removeSteps

	if removePoint <= 0 {
		return head
	}

	for i := range removePoint {
		if i == (removePoint - 1) {
			currentNode.nextNode = currentNode.nextNode.nextNode
			break
		}

		currentNode = currentNode.nextNode
	}

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
		currentNode = currentNode.nextNode
	}

	return count
}

func (list *LinkedList) AppendNode(data int) {
	newNode := &Node{value: data}

	if list.FirstNode == nil {
		list.FirstNode = newNode
		return
	}

	currentNode := list.FirstNode

	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	currentNode.nextNode = newNode
}

func (list LinkedList) ShowAll() {
	var currentNode = list.FirstNode

	if currentNode == nil {
		fmt.Printf("\nA lista está vázia. \n")
		return
	}

	fmt.Printf("\nLista de elementos : \n")
	fmt.Printf("[")

	for currentNode != nil {
		if currentNode.nextNode == nil {
			fmt.Printf("%d", currentNode.value)
			break
		}

		fmt.Printf("%d, ", currentNode.value)
		currentNode = currentNode.nextNode
	}

	fmt.Printf("]")
}
