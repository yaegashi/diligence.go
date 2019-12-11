package main

import "fmt"

func calc(d []int, dSum int) []int {
	dp := make([][]int, len(d)+1)
	for i := range dp {
		dp[i] = make([]int, dSum+1)
	}
	dp[0][0] = 1
	for i := 0; i <= dSum; i++ {
		for j := 1; j <= len(d); j++ {
			dp[j][i] = dp[j-1][i]
			if i >= d[j-1] {
				dp[j][i] += dp[j-1][i-d[j-1]]
			}
		}
	}
	return dp[len(d)]
}

func main() {
	var N, A int
	fmt.Scan(&N, &A)
	var aSum, bSum int
	var a, b []int
	var e uint
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		if x > A {
			a = append(a, x-A)
			aSum += x - A
		} else if x < A {
			b = append(b, A-x)
			bSum += A - x
		} else {
			e++
		}
	}
	aComb := calc(a, aSum)
	bComb := calc(b, bSum)
	sum := 0
	for i := 1; i < len(aComb) && i < len(bComb); i++ {
		sum += aComb[i] * bComb[i]
	}
	sum <<= e
	sum += (1 << e) - 1
	fmt.Println(sum)
}
