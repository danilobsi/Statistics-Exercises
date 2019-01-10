package main

import (
	"fmt"
	"math"
)

func main() {
	var qtty int
	fmt.Scanf("%d\n", &qtty)
	array := make([]int, qtty)

	for i := 0; i < qtty; i++ {
		fmt.Scanf("%d", &array[i])
	}

	fmt.Printf("%.1f", standardDeviation(array))
}

func mean(numbers []int) float64 {
	total := 0
	for _, i := range numbers {
		total += i
	}

	return float64(total) / float64(len(numbers))
}

func standardDeviation(numbers []int) float64 {
	mean := mean(numbers)
	var total float64

	for i := 0; i < len(numbers); i++ {
		total += math.Pow(float64(numbers[i])-mean, float64(2))
	}

	return math.Sqrt(total / float64(len(numbers)))
}
