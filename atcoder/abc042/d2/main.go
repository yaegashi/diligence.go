package main

import "fmt"

const mod = 1e9 + 7

var fact, invFact []int

func powmod(a, m int) int {
	b := 1
	for m > 0 {
		if m&1 != 0 {
			b = b * a % mod
		}
		a = a * a % mod
		m >>= 1
	}
	return b
}

func combmod(n, r int) int {
	c := fact[n]
	c *= invFact[n-r]
	c %= mod
	c *= invFact[r]
	c %= mod
	return c
}

func main() {
	var H, W, A, B int
	fmt.Scan(&H, &W, &A, &B)

	max := H + W + 1
	fact = make([]int, max+1)
	invFact = make([]int, max+1)
	fact[0] = 1
	invFact[0] = 1
	for i := 1; i <= max; i++ {
		fact[i] = (i * fact[i-1]) % mod
		invFact[i] = powmod(i, mod-2)
		invFact[i] = (invFact[i] * invFact[i-1]) % mod
	}

	s := 0
	for i := 0; i < H-A; i++ {
		x := combmod(i+1+B-2, i)
		y := combmod(W-B+H-i-2, H-i-1)
		z := (x * y) % mod
		s = (s + z) % mod
	}
	fmt.Println(s)
}
