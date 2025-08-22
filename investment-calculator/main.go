package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years:")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted with inflation): %.1f\n", futureRealValue)

	// outputs information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf(`Future Value: %.2f]n
	// Future Value (adjusted with inflation): %.2f`, futureValue, futureRealValue`)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

// Function examples that acts like a wrapper utility around the Print function
// func outputText(texts ...string) {
// 	for _, text := range texts {
// 		fmt.Print(text)
// 	}
// }

// func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }