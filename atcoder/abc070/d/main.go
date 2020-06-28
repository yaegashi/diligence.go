package main

// Solution for https://atcoder.jp/contests/abc070/tasks/abc070_d

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

type edge struct{ to, cost int }

func (con *contest) main() error {
	var n int
	con.Scan(&n)
	e := make([][]edge, n)
	for i := 1; i < n; i++ {
		var a, b, c int
		con.Scan(&a, &b, &c)
		a--
		b--
		e[a] = append(e[a], edge{b, c})
		e[b] = append(e[b], edge{a, c})
	}
	var q, k int
	con.Scan(&q, &k)
	c := make([]int, n)
	var dfs func(p, q, r int)
	dfs = func(p, q, r int) {
		c[p] = r
		for _, x := range e[p] {
			if x.to == q {
				continue
			}
			dfs(x.to, p, r+x.cost)
		}
	}
	dfs(k-1, -1, 0)
	for i := 0; i < q; i++ {
		var x, y int
		con.Scan(&x, &y)
		con.Println(c[x-1] + c[y-1])
	}
	return nil
}
