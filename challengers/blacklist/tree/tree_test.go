package tree_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ferreira-gn/estrutura-de-dados/challengers/blacklist/tree"
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

	values := []int32{
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

	values := []int32{
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

	values := []int32{
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

func Test_CheckTheNumberOfSheetsInTheTree(t *testing.T) {
	tree := tree.New()
	sheets := tree.NumberOfSheets()

	if sheets != 0 {
		t.Fatalf("Ao criar uma nova árvore, a quantidade de folhas inicialmente deve ser 0.\nQuantidade de folhas recebida : %v \nQuantidade de folhas esperada : 0", sheets)
	}

	values := []int32{
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

func Test_RemoveLeafNode(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 30} {
		tree.AddNode(value)
	}

	tree.Remove(10)

	if tree.Contains(10) {
		t.Fatalf("Ao remover um nó folha, ele não deve continuar presente na árvore.")
	}

	if tree.Size() != 2 {
		t.Fatalf("Ao remover um nó folha, o tamanho da árvore deve diminuir.\nTamanho recebido : %v \nTamanho esperado : 2", tree.Size())
	}
}

func Test_RemoveNodeWithOneChild(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 5, 30} {
		tree.AddNode(value)
	}

	tree.Remove(10)

	if tree.Contains(10) {
		t.Fatalf("Ao remover um nó com um filho, ele não deve continuar presente na árvore.")
	}

	if !tree.Contains(5) {
		t.Fatalf("Ao remover um nó com um filho, o filho deve continuar presente na árvore.")
	}
}

func Test_RemoveNodeWithTwoChildren(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{50, 30, 70, 20, 40, 60, 80} {
		tree.AddNode(value)
	}

	tree.Remove(50)

	if tree.Contains(50) {
		t.Fatalf("Ao remover um nó com dois filhos, ele não deve continuar presente na árvore.")
	}

	if !tree.Contains(60) {
		t.Fatalf("Ao remover um nó com dois filhos, o sucessor deve continuar presente ocupando a posição removida.")
	}

	if tree.Size() != 6 {
		t.Fatalf("Ao remover um nó com dois filhos, o tamanho da árvore deve diminuir.\nTamanho recebido : %v \nTamanho esperado : 6", tree.Size())
	}
}

func Test_RemoveValueThatDoesNotExist(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 30} {
		tree.AddNode(value)
	}

	tree.Remove(99)

	if tree.Size() != 3 {
		t.Fatalf("Ao remover um valor que não existe, o tamanho da árvore não deve mudar.\nTamanho recebido : %v \nTamanho esperado : 3", tree.Size())
	}

	if !tree.Contains(20) || !tree.Contains(10) || !tree.Contains(30) {
		t.Fatalf("Ao remover um valor que não existe, os nós existentes devem continuar presentes.")
	}
}

func Test_SearchInEmptyTree(t *testing.T) {
	tree := tree.New()

	if tree.Search(10) != nil {
		t.Fatalf("Ao buscar em uma árvore vazia, deve retornar nil.")
	}

	if tree.SearchInOrder(10) != nil {
		t.Fatalf("Ao buscar em ordem em uma árvore vazia, deve retornar nil.")
	}

	if tree.Contains(10) {
		t.Fatalf("Ao verificar um valor em uma árvore vazia, deve retornar false.")
	}
}

func Test_SearchInOrderFindsExistingNode(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 30, 5, 15} {
		tree.AddNode(value)
	}

	foundNode := tree.SearchInOrder(15)

	if foundNode == nil {
		t.Fatalf("Ao buscar em ordem um valor existente, deve retornar o nó encontrado.")
	}
}

func Test_ReportInOrderWithEmptyTree(t *testing.T) {
	tree := tree.New()

	output := captureOutput(func() {
		tree.ReportInOrder()
	})

	if !strings.Contains(output, "RELATORIO DE IPS BLOQUEADOS") {
		t.Fatalf("O relatório deve conter o título.\nSaída recebida:\n%s", output)
	}

	if !strings.Contains(output, "Nenhum IP bloqueado.") {
		t.Fatalf("O relatório de árvore vazia deve informar que não há IPs bloqueados.\nSaída recebida:\n%s", output)
	}
}

func Test_ReportInOrderShowsIPsInAscendingOrder(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 30, 5, 15} {
		tree.AddNode(value)
	}

	output := captureOutput(func() {
		tree.ReportInOrder()
	})

	expectedParts := []string{
		"RELATORIO DE IPS BLOQUEADOS",
		"IP              TENTATIVAS",
		"5               1",
		"10              1",
		"15              1",
		"20              1",
		"30              1",
	}

	for _, expectedPart := range expectedParts {
		if !strings.Contains(output, expectedPart) {
			t.Fatalf("O relatório deve conter %q.\nSaída recebida:\n%s", expectedPart, output)
		}
	}

	assertTextOrder(t, output, "5", "10", "15", "20", "30")
}

func Test_ReportInOrderShowsAttemptsForDuplicatedIP(t *testing.T) {
	tree := tree.New()

	for _, value := range []int32{20, 10, 20, 20, 30} {
		tree.AddNode(value)
	}

	output := captureOutput(func() {
		tree.ReportInOrder()
	})

	if !strings.Contains(output, "20              3") {
		t.Fatalf("O relatório deve exibir a quantidade de tentativas para IPs repetidos.\nSaída recebida:\n%s", output)
	}
}

func captureOutput(fn func()) string {
	oldStdout := os.Stdout
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	fn()
	writer.Close()
	os.Stdout = oldStdout

	var buffer bytes.Buffer
	io.Copy(&buffer, reader)
	return buffer.String()
}

func assertTextOrder(t *testing.T, text string, values ...string) {
	t.Helper()

	lastIndex := -1
	for _, value := range values {
		currentIndex := strings.Index(text, value)
		if currentIndex == -1 {
			t.Fatalf("Valor %q não encontrado no texto.\nTexto recebido:\n%s", value, text)
		}

		if currentIndex < lastIndex {
			t.Fatalf("Valor %q apareceu fora de ordem.\nTexto recebido:\n%s", value, text)
		}

		lastIndex = currentIndex
	}
}
