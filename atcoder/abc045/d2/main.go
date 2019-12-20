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
	r := make([]int, 0, len(M))
	for i, c := range M {
		sort.Ints(c)
		r = append(r, i)
	}
	sort.Ints(r)
	A := make([]int, 10)
	h := 0
	v := 0
	q := make([][]int, 3)
	for h = 0; h < 2; h++ {
		if len(r) > 0 && h == r[0] {
			q[h+1] = M[h]
			r = r[1:]
		}
	}
	for {
		q[0], q[1], q[2] = q[1], q[2], nil
		if q[0] == nil && q[1] == nil && q[2] == nil {
			h = H
			if len(r) > 0 {
				h = r[0]
			}
		}
		if h >= H {
			break
		}
		if len(r) > 0 && h == r[0] {
			q[2] = M[h]
			r = r[1:]
		}
		calc(W, A, [][]int{q[0], q[1], q[2]})
		h++
		v++
	}
	A[0] += (H - 2 - v) * (W - 2)
	for _, a := range A {
		fmt.Println(a)
	}
}
