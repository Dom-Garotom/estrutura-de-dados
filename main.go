package main

import (
	"fmt"
	moduleLinkedList "helloWOrld/linkedList"
)

func main() {
	var linkedList = moduleLinkedList.LinkedList{}

	linkedList.ShowAll()

	fmt.Printf("\nadicionando valores....\n")

	linkedList.AppendNode("Node 01")
	linkedList.AppendNode("Node 02")
	linkedList.AppendNode("Node 03")
	linkedList.AppendNode("Node 04")

	linkedList.ShowAll()
	fmt.Printf("\nTamanho final da lista : %d\n", linkedList.Size())

	fmt.Printf("\nProcurando por Node 01....\n")

	node := linkedList.FindNode("Node 01")

	if node == nil {
		fmt.Printf("\n\nNode 01 não foi encontrado.\n")
		return
	}

	fmt.Printf("Elemento encontrado ")

	fmt.Printf("\n\nVerificando se um ponteiro aleátorio existe em nossa lista.....\n\n")

	var ptr = moduleLinkedList.NewNode("Novo node")
	var exist bool = linkedList.Contains(ptr)

	if !exist {
		fmt.Printf("\nPonteiro '%s' não existe na lista encadeada", linkedList.GetValue(ptr))
	}

}
