package main

import (
	"fmt"
	"sort"
)

func main() {
	i := make([]int, 3)
	fmt.Scan(&i[0], &i[1], &i[2])
	sort.Ints(i)
	if i[0] == 5 && i[1] == 5 && i[2] == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
