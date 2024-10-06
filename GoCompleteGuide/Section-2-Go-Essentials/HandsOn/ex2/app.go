package main

import (
	"fmt"
	"math"
)

func main() {
	// Some variables
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	excpectedReturRate := 5.5

	// User input
	fmt.Println("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// User input
	fmt.Println("Expected Return: ")
	fmt.Scan(&excpectedReturRate)

	// User input
	fmt.Println("Years: ")
	fmt.Scan(&years)

	// Some calculations.
	featureValue := investmentAmount * math.Pow(1+excpectedReturRate/100, years)
	realFeatureValue := featureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(featureValue)
	fmt.Println(realFeatureValue)

}
