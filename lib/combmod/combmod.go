package main

const mod = 1e9 + 7

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
	c := 1
	for i := 0; i < r; i++ {
		c *= n - i
		c %= mod
		c *= powmod(i+1, mod-2)
		c %= mod
	}
	return c
}
