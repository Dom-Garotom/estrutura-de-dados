package main

import (
	"bufio"
	"fmt"
	"os"

	expressionbalanced "github.com/Dom-Garotom/BalanceamentoDeExpressoes/isExpressionBalanced"
)

/*
Main ->  Contém o código de exibição que você está vendo.
Data -> Contém o arquivo de texto com as expressões a serem verificadas.
isExpressionBalanced -> Contém o código da função que faz a verificação do balanceamento de uma expressõa.
Satck -> contém o código da implementação da pilha usada para resolver o desáfio.
*/

func main() {
	file, error := os.Open("data/expression.txt")

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
