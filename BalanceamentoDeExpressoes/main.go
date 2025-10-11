package main

import (
	"fmt"

	"github.com/Dom-Garotom/BalanceamentoDeExpressoes/stack"
)

func main() {

	// recebemos um expressão

	const exp string = "{[()]}"
	const openCharacter string = "{[("
	const closeCharacter string = "}])"

	var stack = stack.CreateStack()
	fmt.Print(stack.IsEmpty())
	// salvamos os valores de abertura em uma pilha

	// verificamos se os valores de abertura e fechamento batem

	// retornamos os resultados
}
