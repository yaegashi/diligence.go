package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := -1
	b := -1
	for j := 0; j < len(s); j++ {
		if j < len(s)-1 && s[j] == s[j+1] {
			a = j + 1
			b = j + 2
			break
		}
		if j < len(s)-2 && s[j] == s[j+2] {
			a = j + 1
			b = j + 3
			break
		}
	}
	fmt.Println(a, b)
}
