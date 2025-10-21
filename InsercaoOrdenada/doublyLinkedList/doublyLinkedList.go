package doublylinkedlist

import "fmt"

type DoublyLinkedList struct {
	head *Node
}

func (list *DoublyLinkedList) AppendNode(value int) {
	newNode := Node{value: value}
	currentNode := list.head

	if list.head == nil {
		list.head = &newNode
		return
	}

	for currentNode != nil {
		if currentNode.nextNode == nil {
			currentNode.nextNode = &newNode
			newNode.previousNode = currentNode
			break
		}

		if newNode.value >= currentNode.value && newNode.value < currentNode.nextNode.value {
			newNode.nextNode = currentNode.nextNode
			newNode.previousNode = currentNode

			currentNode.nextNode.previousNode = &newNode
			currentNode.nextNode = &newNode
			break
		}

		currentNode = currentNode.nextNode
	}
}

func (list DoublyLinkedList) ShowAll() {
	currentNode := list.head

	if currentNode == nil {
		fmt.Printf("A lista está vazia")
		return
	}

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
