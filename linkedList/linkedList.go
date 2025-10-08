/*
Package linkedlist fornece estruturas e funções para manipulação
de listas encadeadas simples.
*/
package linkedlist

import "fmt"

type LinkedList struct {
	FirstNode *Node
}

/*
AppendNode é uma função que adiciona um novo nô dentro da nossa lista encadeada.

	Recebe uma string como parâmetro
*/
func (list *LinkedList) AppendNode(data string) {
	var newNode *Node = &Node{data: data}

	if list.FirstNode == nil {
		list.FirstNode = newNode
		return
	}

	var currentNode *Node = list.FirstNode

	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	currentNode.nextNode = newNode
}

/*
ShowAll é uma função responsável por apresentar todos os elementos da lista
*/
func (list *LinkedList) ShowAll() {
	var currentNode = list.FirstNode

	if currentNode == nil {
		fmt.Printf("\nA lista está vázia. \n")
		return
	}

	fmt.Printf("\nLista de elementos : \n")

	for currentNode != nil {
		fmt.Printf("Elemento - %s\n", currentNode.data)
		currentNode = currentNode.nextNode
	}
}

/*
Size é uma função responsável por retornar o tamanho atual da lista.
*/
func (list *LinkedList) Size() int {
	var currentNode = list.FirstNode
	var count int

	if currentNode == nil {
		return 0
	}

	for currentNode != nil {
		count++
		currentNode = currentNode.nextNode
	}

	return count
}

/*
FindNode é uma função responsável por procurar um elemento específo dentro de uma lista e retornar esse elemento.
*/
func (list *LinkedList) FindNode(findValue string) *Node {
	var currentNode = list.FirstNode

	if findValue == "" || currentNode == nil {
		return nil
	}

	for currentNode != nil {
		if currentNode.data == findValue {
			return currentNode
		}

		currentNode = currentNode.nextNode
	}

	return nil
}

/*
Contains é uma função responsável por informar se um elemento existe em uma lista e retornar um valor booleano correspondente
*/
func (list *LinkedList) Contains(node *Node) bool {
	var currentNode = list.FirstNode

	if currentNode == nil || node == nil {
		return false
	}

	for currentNode != nil {
		if &currentNode == &node {
			return true
		}

		currentNode = currentNode.nextNode
	}

	return false
}

/*
	concluidas :
	(a) Construtor - Inicializa a classe
	(i) Inserir - Insere um elemento na lista
	(g) mostrarALL - Retorna todos os elementos da lista
	(e) Tamanho - Retorna o tamanho da lista
	(f) Existe - Retorna se um n ́o existe na lista
	(h) Buscar - Retorna se o n ́o cont ́em na lista


	restantes a serem feitas :

(b) ObterProximo - Recebe como argumento um n ́o e retorna o pr ́oximo
(c) ObterValor - Recebe como argumento um n ́o e retorna os valores armazenados dentro
delea
(d) AlterarNo - Recebe como argumento um n ́o e dois interios para alterar as informa ̧c ̃oes
do n ́o referenciado




(j) Excluir - Exclui um elemento da list
(k) Destrutor - Destr ́oi um n ́o
*/
