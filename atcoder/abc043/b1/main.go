package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s)
	for _, c := range s {
		switch c {
		case 'B':
			if len(t) > 0 {
				t = t[:len(t)-1]
			}
		default:
			t += string(c)
		}
	}
	fmt.Println(t)
}
