package main

import (
	"fmt"
	"math"
)

func main() {
	values := []int{1, 10, 5000, -20, 35}
	fmt.Println("Min: ", min(values...))
	fmt.Println("Max: ", max(values...))
}

func min(x ...int) int {
	minValue := math.MaxInt64
	for _, value := range x {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func max(x ...int) int {
	maxValue := math.MinInt64
	for _, value := range x {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
