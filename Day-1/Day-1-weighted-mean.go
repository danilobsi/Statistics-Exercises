package main

import "fmt"

func main() {
	var qtty int
	fmt.Scanf("%d", &qtty)

	weights := make([]int, qtty)
	itens := make([]int, qtty)

	for i := 0; i < qtty; i++ {
		fmt.Scanf("%d", &weights[i])
	}

	for i := 0; i < qtty; i++ {
		fmt.Scanf("%d", &itens[i])
	}

	fmt.Printf("%.1f\n", weightedMean(weights, itens))
}

func weightedMean(weightArray []int, itensArray []int) float64 {
	weightedAmount := 0
	itensSum := 0
	for i := 0; i < len(weightArray); i++ {
		weightedAmount += weightArray[i] * itensArray[i]
		itensSum += itensArray[i]
	}

	return float64(weightedAmount) / float64(itensSum)
}
