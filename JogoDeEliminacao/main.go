package main

import (
	"fmt"
	circularlist "github.com/Dom-Garotom/estruturaDeDados/jogoDeEliminacao/circularList"
)

func main() {

	list := circularlist.CircularList{}

	list.AppendNode("Marcos")
	list.AppendNode("Adrian")
	list.AppendNode("Matheus")
	list.AppendNode("Lucian")
	list.AppendNode("Natan")
	list.AppendNode("Vitor")
	list.AppendNode("Gean")

	fmt.Print("Lista de players : ")
	list.ShowAll()

	list.RemoveNodeInNthSteps(25)

	fmt.Print("\n\nLista final de players : ")
	list.ShowAll()
}
