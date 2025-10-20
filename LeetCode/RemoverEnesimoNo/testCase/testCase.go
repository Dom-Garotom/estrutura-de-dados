/*
Package testcase é  um pacote responsável por gerenciar os casos de testes do desafio do leetcode
*/
package testcase

import (
	"fmt"
	"reflect"

	linkedlist "github.com/Dom-Garotom/estruturaDeDados/LeetCode/RemoverEnesimoNo/linkedList"
)

func listToSlice(list *linkedlist.LinkedList) []int {
	var result []int
	currentNode := list.FirstNode

	for currentNode != nil {
		result = append(result, currentNode.Value)
		currentNode = currentNode.NextNode
	}

	return result
}

func verifyResult(expectResult, result []int) {
	fmt.Printf("Saída esperada : %v\n", expectResult)
	fmt.Printf("Saída recebida : %v\n", result)

	if reflect.DeepEqual(expectResult, result) {
		fmt.Printf("✅ Teste aprovado!\n\n")
		return
	}

	fmt.Printf("❌ Teste recusado!\n\n")
}

func TestCase01() {
	fmt.Println("\033[1;35mTestCase01\033[m - Remove o 4º nó a partir do final em lista de 10 elementos")
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 10 {
		elementList.AppendNode(i + 1)
	}

	fmt.Printf("Entrada : ")
	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 4)

	expectResult := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	result := listToSlice(&elementList)

	verifyResult(expectResult, result)
}

func TestCase02() {
	fmt.Println("\033[1;35mTestCase02\033[m - Remove o último nó da lista")
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 10 {
		elementList.AppendNode(i + 1)
	}

	fmt.Printf("Entrada : ")
	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 1)

	expectResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := listToSlice(&elementList)

	verifyResult(expectResult, result)
}

func TestCase03() {
	fmt.Println("\033[1;35mTestCase03\033[m - Remove o primeiro nó da lista")
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 5 {
		elementList.AppendNode(i + 1)
	}

	fmt.Printf("Entrada : ")
	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 5)

	expectResult := []int{2, 3, 4, 5}
	result := listToSlice(&elementList)

	verifyResult(expectResult, result)
}

func TestCase04() {
	fmt.Println("\033[1;35mTestCase04\033[m - Lista com 1 elemento, remove o único nó")
	elementList := linkedlist.LinkedList{FirstNode: nil}

	elementList.AppendNode(42)

	fmt.Printf("Entrada : ")
	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 1)

	expectResult := []int{}
	result := listToSlice(&elementList)

	verifyResult(expectResult, result)
}

func TestCase05() {
	fmt.Println("\033[1;35mTestCase05\033[m - removeSteps maior que o tamanho da lista")
	elementList := linkedlist.LinkedList{FirstNode: nil}

	for i := range 3 {
		elementList.AppendNode(i + 1)
	}

	fmt.Printf("Entrada : ")
	elementList.ShowAll()

	linkedlist.RemoveNthNodeFromEndList(&elementList, 10)

	expectResult := []int{1, 2, 3}
	result := listToSlice(&elementList)

	verifyResult(expectResult, result)
}
