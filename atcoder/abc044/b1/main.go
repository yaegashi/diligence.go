package main

import "fmt"

func main() {
	cMap := map[rune]int{}
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		cMap[c]++
	}
	output := "Yes"
	for _, c := range cMap {
		if c&1 != 0 {
			output = "No"
			break
		}
	}
	fmt.Println(output)
}
