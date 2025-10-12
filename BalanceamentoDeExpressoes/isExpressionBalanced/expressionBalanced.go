/*Package expressionbalanced é um pacote que contém uma função que verifica o balanceamento de expressões matemáticas*/
package expressionbalanced

import (
	"strings"

	"github.com/Dom-Garotom/BalanceamentoDeExpressoes/stack"
)

func IsExpressionBalanced(exp string) bool {
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
	}

	return expressionStack.IsEmpty()
}
