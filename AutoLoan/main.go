package main

import (
	"fmt"
	"math"
)

func main() {
	var price, monthlyPayment, annualRate float64
	var loanTerm int

	fmt.Scanln(&price)
	fmt.Scanln(&monthlyPayment)
	fmt.Scanln(&loanTerm)

	var annualRateU float64 = 1.0
	var annualRateL float64 = 0.0

	for {
		if (annualRateU/annualRateL-1.0) < math.Pow10(-9) && (annualRateU-annualRateL)/annualRateL < math.Pow10(-9) {
			break
		}
		balance := price
		annualRate = (annualRateU + annualRateL) / 2.0
		for i := 0; i < loanTerm; i++ {
			balance = balance*(1.0+annualRate) - monthlyPayment
		}

		if balance > 0 {
			annualRateU = annualRate
		} else {
			annualRateL = annualRate
		}
	}

	fmt.Println(annualRate * 12.0 * 100.0)
}
