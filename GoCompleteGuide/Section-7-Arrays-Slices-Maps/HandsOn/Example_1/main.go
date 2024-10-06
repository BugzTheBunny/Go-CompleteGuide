package main

import "fmt"

func main() {
	var products [4]string
	prices := [5]float64{10, 20, 30, 40, 44}
	fmt.Println(prices)
	fmt.Println(products)

	products[0] = "potato"
	products[3] = "banana"

	fmt.Println(products)

	featuredPrices := prices[1:3] // First value included, second value EXCLUDED
	fmt.Println(featuredPrices)

	newPrices := []float64{10.90, 20.99, 30.99}
	fmt.Println(newPrices)

	newPrices[1] = 9
	//newPrices[3] = 20 // This will not work, as this index does not exist.

	newPrices = append(newPrices, 9)
	fmt.Println(newPrices)

	pricesToMerge := []float64{100, 200, 300, 400}
	newPrices = append(newPrices, pricesToMerge...)
	fmt.Println(newPrices)

}
