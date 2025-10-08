package main

import (
	"fmt"
	moduleLinkedList "helloWOrld/linkedList"
)

func main() {
	var linkedList = moduleLinkedList.LinkedList{}

	fmt.Printf("\nadicionando valores....\n")

	linkedList.AppendNode("Node 01")
	linkedList.AppendNode("Node 02")
	linkedList.AppendNode("Node 03")
	linkedList.AppendNode("Node 04")

	linkedList.ShowAll()

	fmt.Printf("\nAtualizando os valores....\n")

	secondNode := linkedList.FindNode("Node 02")
	linkedList.UpdateNode(secondNode, "New Node 02")

	linkedList.ShowAll()

}
