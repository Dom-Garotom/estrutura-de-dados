package main

import (
	"fmt"
	linkedlist "helloWOrld/linkedList"
)

// Informando aqui o import acima da linkedList é o modulo criado por mim da lista encadeada segundo a questão 01 e 02 da atividade de listas encadeadas.

func main() {
	list := linkedlist.LinkedList{}
	fmt.Println("Lista encadeada inicializada com sucesso!")

	list.AppendNode("A")
	list.AppendNode("B")
	list.AppendNode("C")
	list.AppendNode("D")
	fmt.Println("\n--- Elementos adicionados ---")
	list.ShowAll()

	fmt.Printf("\nTamanho atual da lista: %d\n", list.Size())

	fmt.Println("\n--- Buscando o elemento 'C' ---")
	node := list.FindNode("C")
	if node != nil {
		fmt.Printf("Elemento encontrado: %s\n", list.GetValue(node))
	} else {
		fmt.Println("Elemento não encontrado.")
	}

	fmt.Println("\n--- Verificando se o nó 'C' existe na lista ---")
	fmt.Println("Existe:", list.Contains(node))

	fmt.Println("\n--- Obtendo valor do nó encontrado ---")
	fmt.Printf("Valor armazenado: %s\n", list.GetValue(node))

	fmt.Println("\n--- Obtendo próximo nó de 'C' ---")
	nextNode := list.GetNextNode(node)
	if nextNode != nil {
		fmt.Printf("Próximo elemento: %s\n", list.GetValue(nextNode))
	} else {
		fmt.Println("Não há próximo nó.")
	}

	fmt.Println("\n--- Atualizando o valor do nó 'B' para 'Beta' ---")
	nodeB := list.FindNode("B")
	list.UpdateNode(nodeB, "Beta")
	list.ShowAll()

	fmt.Println("\n--- Removendo o elemento 'C' ---")
	list.DeleteNode(node)
	list.ShowAll()
	fmt.Printf("Tamanho atual da lista: %d\n", list.Size())

	fmt.Println("\n--- Destruindo a lista ---")
	list.Destroy()
	list.ShowAll()
}
