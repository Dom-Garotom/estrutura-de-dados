/*
Package testcase é um pacote responsável por gerenciar os casos de testes do desafio do leetcode
*/
package testcase

import linkedlist "github.com/Dom-Garotom/estruturaDeDados/LeetCode/RemoverEnesimoNo/linkedList"

func TestCase01() {
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 10 {
		elementList.AppendNode(i)
	}

	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 4)

	elementList.ShowAll()
}

func TestCase02() {
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 10 {
		elementList.AppendNode(i)
	}

	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 4)

	elementList.ShowAll()
}
