package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Printf("%.1f\n", mean(a))
	fmt.Printf("%.1f\n", median(a))
	fmt.Println(mode(a))
}

func mean(numbers []int) float64 {
	total := 0
	for _, i := range numbers {
		total += i
	}

	return float64(total) / float64(len(numbers))
}

func median(number []int) float64 {
	sort.Ints(number)

	middle := len(number) / 2

	if math.Mod(float64(len(number)), float64(2)) == 1 {
		return float64(number[middle])
	}

	number1 := float64(number[middle])
	number2 := float64(number[middle-1])

	return (number1 + number2) / 2
}

func mode(numbers []int) int {
	sort.Ints(numbers)

	modeArray := [][]int{}

	for _, i := range numbers {
		found := false

		for _, modeItem := range modeArray {
			if modeItem[0] == i {
				modeItem[1]++
				found = true
				break
			}
		}

		if !found {
			modeArray = append(modeArray, []int{i, 1})
		}
	}

	var modeNumber = 0
	var modeQtty = 0
	for _, i := range modeArray {
		if i[1] > modeQtty {
			modeNumber = i[0]
			modeQtty = i[1]
		}
	}

	return modeNumber
}

func convertToIntArray(array []string) []int {
	intArray := []int{}
	for _, i := range array {
		intArray = append(intArray, convertToInt(i))
	}

	return intArray
}

func convertToInt(value string) int {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}
	return intValue
}
