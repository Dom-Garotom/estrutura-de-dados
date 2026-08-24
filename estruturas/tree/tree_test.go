package tree_test

import (
	"testing"

	"github.com/ferreira-gn/estrutura-de-dados/estruturas/tree"
)

func Test_CreateNewTree(t *testing.T) {
	tree := tree.New()

	if tree == nil {
		t.Fatalf("Ao criar uma nova árvore binária deve ser retornado um ponteiro válido para a estrutura.\nRecebido : %v", tree)
	}

	if tree.IsEmpty() == false {
		t.Fatalf("Ao criar uma nova árvore espera-se que ela estejá vázia, a arvore contém elementos. \nValor recebido de IsEmpty : %v \nValor esperado de IsEmpty : true", tree.IsEmpty())
	}
}

func Test_AddNewNodeIntoTheTree(t *testing.T) {
	tree := tree.New()

	values := []int64{
		20, 30, 15, 10, 5, 18, 25, 40, 35, 50, 17,
	}

	for _, value := range values {
		tree.AddNode(value)
	}

	if tree.IsEmpty() {
		t.Fatalf("Ao adicinar novos nós em uma árvore, espera-se que ela não continue vázia. \nValor recebido de IsEmpty : %v \nValor esperado de IsEmpty : false ", tree.IsEmpty())
	}

	foundNode := tree.Search(10)

	if foundNode == nil {
		t.Fatalf("Ao adicionar um novo nó em uma árvore deve ser possível encontra-lo.\nRecebido : %v \nEsperado : *Node", foundNode)
	}
}

func Test_CheckTheHeightOfTree(t *testing.T) {
	tree := tree.New()
	treeHeight := tree.Height()

	if treeHeight != 0 {
		t.Fatalf("Ao criar uma nova árvore, a sua altura inicial deve ser 0.\nAltura recebida : %v \nAltura esperada : 0", treeHeight)
	}

	values := []int64{
		20, 30, 15, 10, 5, 18, 25, 40, 35, 50, 17,
	}

	for _, value := range values {
		tree.AddNode(value)
	}

	treeHeight = tree.Height()

	if treeHeight != 4 {
		t.Fatalf("Calculo da altura da árvore está incorreta.\nAltura recebida : %v \nAltura esperada : 4", treeHeight)
	}
}

func Test_CheckTheSizeOfTree(t *testing.T) {
	tree := tree.New()
	treeSize := tree.Size()

	if treeSize != 0 {
		t.Fatalf("Ao criar uma nova árvore, o seu tamanho (número de nodes) inicial deve ser 0.\nTamanho recebida : %v \nTamanho esperada : 0", treeSize)
	}

	values := []int64{
		20, 30, 15, 10, 5, 18, 25, 40, 35, 50, 17,
	}

	for _, value := range values {
		tree.AddNode(value)
	}

	treeSize = tree.Size()

	if treeSize != 11 {
		t.Fatalf("Calculo da tamanho da árvore está incorreta.\nTamanho recebida : %v \nTamanho esperada : 11", treeSize)
	}
}

func Test_CheckTheNumberOfSheetsInTheTree(t *testing.T){
	tree := tree.New()
	sheets := tree.NumberOfSheets()

	if sheets != 0 {
		t.Fatalf("Ao criar uma nova árvore, a quantidade de folhas inicialmente deve ser 0.\nQuantidade de folhas recebida : %v \nQuantidade de folhas esperada : 0", sheets)
	}

	values := []int64{
		20, 30, 15, 10, 5, 18, 25, 40, 35, 50, 17,
	}

	for _, value := range values {
		tree.AddNode(value)
	}

	sheets = tree.NumberOfSheets()

	if sheets != 5 {
		t.Fatalf("Quantidade de folhas da árvore está incorreta.\nQuantidade de folhas recebida : %v \nQuantidade de folhas esperada : 11", sheets)
	}
}
