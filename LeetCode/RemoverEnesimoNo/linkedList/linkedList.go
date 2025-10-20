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
			currentNode.NextNode = currentNode.NextNode.NextNode
			break
		}

		currentNode = currentNode.NextNode
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
