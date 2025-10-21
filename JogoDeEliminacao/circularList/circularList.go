/*
Package circularlist fornece estruturas e funções para manipulação
de listas  circulares encadeadas.
*/
package circularlist

import "fmt"

type CircularList struct {
	head *Node
	end  *Node
}

func (list *CircularList) AppendNode(namePerson string) {
	newNode := &Node{name: namePerson}

	if list.head == nil {
		list.head = newNode
		list.end = newNode
		newNode.nextNode = list.head
		return
	}

	list.end.nextNode = newNode
	list.end = newNode
	newNode.nextNode = list.head
}

func (list *CircularList) RemoveNodeInNthSteps(steps int) {

	if list.head == nil {
		fmt.Printf("\nA lista está vazia")
		return
	}

	if list.end == list.head {
		fmt.Printf("\nA lista contem apenas um único player")
		fmt.Printf("\n[%s]", list.head.name)
		return
	}

	currentNode := list.head
	var previousNode *Node

	for list.head.nextNode != list.head {

		for range steps {
			previousNode = currentNode
			currentNode = currentNode.nextNode
		}

		fmt.Printf("\nRemovendo %s", currentNode.name)

		previousNode.nextNode = currentNode.nextNode

		if currentNode == list.head {
			list.head = currentNode.nextNode
		}

		if currentNode == list.end {
			list.end = previousNode
		}

		currentNode = previousNode.nextNode
	}

	fmt.Printf("%s \n\n", list.end.name)

	fmt.Printf("\nGanhador do jogo")
	fmt.Printf("\n[%s]", list.head.name)

}

func (list CircularList) ShowAll() {
	if list.head == nil {
		fmt.Printf("\nA lista está vazia")
		return
	}

	fmt.Printf("\n[")
	currentNode := list.head

	for {
		if currentNode.nextNode == list.head {
			fmt.Printf("%s", currentNode.name)
			break
		}

		fmt.Printf("%s, ", currentNode.name)
		currentNode = currentNode.nextNode

		if currentNode == list.head {
			break
		}
	}

	fmt.Printf("]\n")
}
