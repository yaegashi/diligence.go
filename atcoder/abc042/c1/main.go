package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	d := make([]int, k)
	var e [10]bool
	var f []int
	for i := range d {
		fmt.Scan(&d[i])
		e[d[i]] = true
	}
	for i := 0; i < 10; i++ {
		if !e[i] {
			f = append(f, i)
		}
	}
	var x []int
	for n > 0 {
		x = append(x, n%10)
		n /= 10
	}
	p := len(x) - 1
	for p >= 0 {
		if e[x[p]] {
			break
		}
		p--
	}
	if p >= 0 {
		for {
			c := false
			j := x[p]
			for {
				j++
				if j == 10 {
					j = 0
					c = true
				}
				if !e[j] {
					break
				}
			}
			x[p] = j
			if !c {
				break
			}
			p++
			if p == len(x) {
				if f[0] == 0 {
					x = append(x, f[1])
				} else {
					x = append(x, f[0])
				}
				break
			}
		}
		for i := p - 1; i >= 0; i-- {
			x[i] = f[0]
		}
	}
	for i := len(x) - 1; i >= 0; i-- {
		fmt.Print(x[i])
	}
}
