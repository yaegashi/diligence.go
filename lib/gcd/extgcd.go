// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm

package main

func extgcd(a, b int, x, y *int) int {
	d := a
	if b > 0 {
		d = extgcd(b, a%b, y, x)
		*y -= a / b * *x
	} else {
		*x = 1
		*y = 0
	}
	return d
}
