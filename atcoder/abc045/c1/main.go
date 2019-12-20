package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	a := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			k, _ := strconv.Atoi(s[i:j])
			b := 0
			if i > 0 {
				b += i - 1
			}
			if j < len(s) {
				b += len(s) - j - 1
			}
			k <<= uint(b)
			a += k
		}
	}
	fmt.Println(a)
}
