package main

import (
	//"fmt"
	"fmt"
	"strings"

	"github.com/Dom-Garotom/BalanceamentoDeExpressoes/stack"
)

func isExpressionBalanced() bool {

	const exp string = "{[({[}]})]}"

	const openCharacter string = "{[("
	const closeCharacter string = "}])"
	pairs := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	var expressionStack = stack.CreateStack()

	for _, character := range exp {
		if strings.Contains(openCharacter, string(character)) {
			expressionStack.AddToTop(string(character))
		}

		if strings.Contains(closeCharacter, string(character)) {
			if expressionStack.IsEmpty() {
				return false
			}

			currentCharacter := string(character)
			top := expressionStack.ViewTheTop()

			if top != pairs[currentCharacter] {
				return false
			}

			expressionStack.RemoveFromTop()
		}
		expressionStack.ViewStack()
	}

	fmt.Printf("\n\nRun final")
	expressionStack.ViewStack()

	return true
}

func main() {
	result := isExpressionBalanced()
	fmt.Printf("Result ", result)
}
