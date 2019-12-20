package main

import (
	"fmt"
	"sort"
)

func calc(W int, A []int, p [][]int) {
	w := 0
	v := 0
	q := []int{0, 0, 0}
	for w = 0; w < 2; w++ {
		for i := 0; i < 3; i++ {
			if len(p[i]) > 0 && w == p[i][0] {
				q[w+1]++
				p[i] = p[i][1:]
			}
		}
	}
	for {
		q[0], q[1], q[2] = q[1], q[2], 0
		if q[0]+q[1]+q[2] == 0 {
			w = W
			for i := 0; i < 3; i++ {
				if len(p[i]) > 0 && p[i][0] < w {
					w = p[i][0]
				}
			}
		}
		if w >= W {
			break
		}
		for i := 0; i < 3; i++ {
			if len(p[i]) > 0 && w == p[i][0] {
				q[2]++
				p[i] = p[i][1:]
			}
		}
		A[q[0]+q[1]+q[2]]++
		w++
		v++
	}
	A[0] += w - 2 - v
}

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)
	M := map[int][]int{}
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		M[a-1] = append(M[a-1], b-1)
	}
	R := make([]int, 0, len(M))
	C := make([][]int, len(M))
	for i := range M {
		R = append(R, i)
	}
	sort.Ints(R)
	for i, j := range R {
		C[i] = M[j]
		sort.Ints(C[i])
	}
	A := make([]int, 10)
	v := 0
	r := []int{-1, -1, -1}
	for i := 0; ; {
		r[0], r[1], r[2] = r[1], r[2], -1
		if i < len(R) && r[0] >= 0 && R[r[0]]+2 == R[i] {
			r[2] = i
			i++
		}
		if i < len(R) && r[1] >= 0 && R[r[1]]+1 == R[i] {
			r[2] = i
			i++
		}
		if r[0] < 0 && r[1] < 0 && r[2] < 0 {
			if i < len(R) {
				r[2] = i
				i++
			} else {
				break
			}
		}
		if r[0] >= 0 {
			if R[r[0]] > H-3 {
				break
			}
		}
		if r[1] >= 0 {
			if R[r[1]] > H-2 {
				break
			}
			if R[r[1]] < 1 {
				continue
			}
		}
		if r[2] >= 0 {
			if R[r[2]] < 2 {
				continue
			}
		}
		c := make([][]int, 3)
		for j := 0; j < 3; j++ {
			if r[j] < 0 {
				break
			}
			c[j] = C[r[j]]
		}
		calc(W, A, c)
		v++
	}
	A[0] += (H - 2 - v) * (W - 2)
	for _, a := range A {
		fmt.Println(a)
	}
}
