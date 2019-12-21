package main

import "fmt"

type p struct{ r, c int }

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)
	M := map[p]int{}
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		for r := a - 3; r < a; r++ {
			if r < 0 || r > H-3 {
				continue
			}
			for c := b - 3; c < b; c++ {
				if c < 0 || c > W-3 {
					continue
				}
				M[p{r, c}]++
			}
		}
	}
	A := make([]int, 10)
	for _, v := range M {
		A[v]++
	}
	A[0] = (H-2)*(W-2) - len(M)
	for _, a := range A {
		fmt.Println(a)
	}
}
