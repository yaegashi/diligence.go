package main

import "fmt"

func main() {
	var N, K, X, Y int
	fmt.Scan(&N, &K, &X, &Y)
	s := 0
	if N < K {
		s += X * N
	} else {
		s += X * K
	}
	if N > K {
		s += Y * (N - K)
	}
	fmt.Println(s)
}
