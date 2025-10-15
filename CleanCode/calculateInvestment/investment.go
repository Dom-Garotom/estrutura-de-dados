package investment

func CalculateInvestment(initialInvestment, monthlyDeposit, annualInterestRateInDecimal float64) (float64, float64) {
	capital := initialInvestment + (monthlyDeposit * 12)
	rate := 1.0 + annualInterestRateInDecimal
	montant := capital * rate
	performance := montant - capital

	return montant, performance
}
