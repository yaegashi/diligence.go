package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}
	minSum := math.MaxInt64
	for x := -100; x <= 100; x++ {
		sum := 0
		for _, y := range a {
			sum += (x - y) * (x - y)
		}
		if sum < minSum {
			minSum = sum
		}
	}
	fmt.Println(minSum)
}
