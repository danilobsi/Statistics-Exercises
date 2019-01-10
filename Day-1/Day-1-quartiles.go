package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var qtty int

	fmt.Scanf("%d\n", &qtty)

	numbers := make([]int, qtty)

	for i := 0; i < qtty; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	q1, q2, q3 := quartile(numbers)

	fmt.Printf("%d\n%d\n%d\n", q1, q2, q3)
}

func quartile(numbers []int) (int, int, int) {
	sort.Ints(numbers)

	medianQ2 := medianS(numbers)

	medianPos := len(numbers) / 2

	addExtraPos := 0
	if math.Mod(float64(len(numbers)), float64(2)) == 1 {
		addExtraPos = 1
	}

	medianQ1 := medianS(numbers[:medianPos])
	medianQ3 := medianS(numbers[medianPos+addExtraPos:])

	return medianQ1, medianQ2, medianQ3
}

func medianS(incSortedNumbers []int) int {
	middle := len(incSortedNumbers) / 2

	if math.Mod(float64(len(incSortedNumbers)), float64(2)) == 1 {
		return incSortedNumbers[middle]
	}

	number1 := incSortedNumbers[middle]
	number2 := incSortedNumbers[middle-1]

	return (number1 + number2) / 2
}
