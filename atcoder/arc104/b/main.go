package main

// Solution for https://atcoder.jp/contests/arc104/tasks/arc104_b

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type contest struct {
	in  io.Reader
	out io.Writer
}

func (con *contest) Scan(a ...interface{}) (int, error) {
	return fmt.Fscan(con.in, a...)
}
func (con *contest) Scanln(a ...interface{}) (int, error) {
	return fmt.Fscanln(con.in, a...)
}
func (con *contest) Scanf(f string, a ...interface{}) (int, error) {
	return fmt.Fscanf(con.in, f, a...)
}
func (con *contest) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(con.out, a...)
}
func (con *contest) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(con.out, a...)
}
func (con *contest) Printf(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(con.out, f, a...)
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	con := &contest{in: in, out: out}
	con.main()
}

func (con *contest) main() error {
	var n int
	var s string
	con.Scan(&n, &s)
	a := make([]int, n+1)
	t := make([]int, n+1)
	c := make([]int, n+1)
	g := make([]int, n+1)
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'A':
			a[i+1] = 1
		case 'T':
			t[i+1] = 1
		case 'C':
			c[i+1] = 1
		case 'G':
			g[i+1] = 1
		}
		a[i+1] += a[i]
		t[i+1] += t[i]
		c[i+1] += c[i]
		g[i+1] += g[i]
	}
	ok := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			da := a[j] - a[i]
			dt := t[j] - t[i]
			dc := c[j] - c[i]
			dg := g[j] - g[i]
			if da == dt && dc == dg {
				ok++
			}
		}
	}
	con.Println(ok)
	return nil
}
