package main

import "fmt"

func main() {
	p := make([]string, 3)
	fmt.Scan(&p[0], &p[1], &p[2])
	i := 0
	for len(p[i]) > 0 {
		i, p[i] = int(p[i][0])-'a', p[i][1:]
	}
	fmt.Println(string(i + 'A'))
}
