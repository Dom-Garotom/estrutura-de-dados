package main

import (
	"bufio"
	"fmt"
	"os"

	expressionbalanced "github.com/Dom-Garotom/BalanceamentoDeExpressoes/isExpressionBalanced"
)

func main() {
	file, error := os.Open("expression.txt")

	if error != nil {
		fmt.Println("Erro ao abrir o arquivo de expressões")
		return
	}

	defer file.Close()

	fmt.Println("\033[36m\n\nResultado do balanceamento das expressões :\033[0m")
	fmt.Printf("\nExpressões  %29s Resultados ", "")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		result := expressionbalanced.IsExpressionBalanced(line)

		if result {
			fmt.Printf("\n%-40s \033[36m está balanceada\033[0m", line)
		} else {
			fmt.Printf("\n%-40s \033[31m não está balanceada\033[0m", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}
}
