package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

type byX struct{ x, y, z []int }

func (b byX) Len() int           { return len(b.z) }
func (b byX) Less(i, j int) bool { return b.x[b.z[i]] < b.x[b.z[j]] }
func (b byX) Swap(i, j int)      { b.z[i], b.z[j] = b.z[j], b.z[i] }

type byY struct{ x, y, z []int }

func (b byY) Len() int           { return len(b.z) }
func (b byY) Less(i, j int) bool { return b.y[b.z[i]] < b.y[b.z[j]] }
func (b byY) Swap(i, j int)      { b.z[i], b.z[j] = b.z[j], b.z[i] }

type edge struct{ u, v, cost int }

type byE []edge

func (b byE) Len() int           { return len(b) }
func (b byE) Less(i, j int) bool { return b[i].cost < b[j].cost }
func (b byE) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func (con *contest) main() error {
	var n int
	con.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)
	e := make([]edge, 0, n*2)
	for i := 0; i < n; i++ {
		con.Scan(&x[i], &y[i])
		z[i] = i
	}
	sort.Sort(byX{x, y, z})
	for i := 0; i < n-1; i++ {
		c := x[z[i+1]] - x[z[i]]
		e = append(e, edge{z[i], z[i+1], c})
	}
	sort.Sort(byY{x, y, z})
	for i := 0; i < n-1; i++ {
		c := y[z[i+1]] - y[z[i]]
		e = append(e, edge{z[i], z[i+1], c})
	}
	sort.Sort(byE(e))
	g := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = i
	}
	res := 0
	for _, a := range e {
		u := a.u
		for u != g[u] {
			u = g[u]
		}
		v := a.v
		for v != g[v] {
			v = g[v]
		}
		if u != v {
			res += a.cost
			g[u] = v
			u = v
		}
		for i := a.u; i != u; i, g[i] = g[i], u {
		}
		for i := a.v; i != v; i, g[i] = g[i], v {
		}
	}
	con.Println(res)
	return nil
}
