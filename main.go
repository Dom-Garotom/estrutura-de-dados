package main

import (
	"fmt"
	linkedList "helloWOrld/linkedLIst"
)



func main () {
	var linkedList = linkedList.LinkedList{}	
	
	linkedList.ShowAll()

	fmt.Printf("\nadicionando valores....\n")

	linkedList.AppendNode("Node 01")
	linkedList.AppendNode("Node 02")
	linkedList.AppendNode("Node 03")
	linkedList.AppendNode("Node 04")

	linkedList.ShowAll()
	fmt.Printf("\nTamanho final da lista : %d\n", linkedList.Size())
}