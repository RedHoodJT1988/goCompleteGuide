package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// Get user input
	println("Enter Revenue:")
	fmt.Scanln(&revenue)

	println("Enter Expenses:")
	fmt.Scanln(&expenses)

	println("Enter Tax Rate (as a decimal):")
	fmt.Scanln(&taxRate)

	// Calculate profit before tax
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
    ratio := ebt / profit

    // Output the results
    fmt.Println(ebt)
    fmt.Println(profit)
    fmt.Println(ratio)
}