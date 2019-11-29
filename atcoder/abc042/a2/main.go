package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	i := make([]int, 3)
	fmt.Fscan(in, &i[0], &i[1], &i[2])
	sort.Ints(i)
	if i[0] == 5 && i[1] == 5 && i[2] == 7 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
