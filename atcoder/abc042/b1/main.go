package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}
	sort.Strings(s)
	for _, t := range s {
		fmt.Print(t)
	}
}
