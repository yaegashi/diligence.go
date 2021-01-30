package main

const mod = 1e9 + 7

func powMod(a, m int) int {
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

func combMod(n, r int) int {
	c := 1
	for i := 0; i < r; i++ {
		c *= n - i
		c %= mod
		c *= powMod(i+1, mod-2)
		c %= mod
	}
	return c
}
