package main

import doublylinkedlist "github.com/Dom-Garotom/estruturaDeDados/insercaoOrdenada/doublyLinkedList"

func main() {
	list := doublylinkedlist.DoublyLinkedList{}

	list.AppendNode(1)
	list.AppendNode(4)
	list.AppendNode(3)
	list.AppendNode(10)
	list.AppendNode(5)
	list.AppendNode(4)
	list.AppendNode(9)
	list.AppendNode(7)
	list.AppendNode(22)
	list.AppendNode(17)
	list.AppendNode(92)
	list.AppendNode(73)

	list.ShowAll()
}
