package main

import (
	"fmt"
	investment "github.com/Dom-Garotom/CleanCode/calculateInvestment"
)

func main() {
	var initialInvestment float64 = 10000.00
	var monthlyDeposit float64 = 2000.00
	var annualInterestRateInDecimal float64 = 0.1
	var timeInvested int32 = 5

	for i := 0; i < int(timeInvested); i++ {
		// O cálculo do rendimento de investimento do usuário é calculado a juros compostos anuais.
		currentInvestment, performance := investment.CalculateInvestment(initialInvestment, monthlyDeposit, annualInterestRateInDecimal)

		fmt.Printf("\nAno : %d \n", i+1)
		fmt.Printf("Saldo: %20s      Rendimento:\n", "")
		fmt.Printf("R$ %.f %23s", currentInvestment, "")
		fmt.Printf("R$ %.f \n", performance)

		initialInvestment = currentInvestment
	}
}
