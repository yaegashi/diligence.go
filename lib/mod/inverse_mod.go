package main

func inverseMod(a, m int) int {
	b := m
	x, y := 1, 0
	for b > 0 {
		q := a / b
		a, b = b, a-q*b
		x, y = y, x-q*y
	}
	if x < 0 {
		x += m
	}
	return x
}
