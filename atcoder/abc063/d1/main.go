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

type greater []int

func (g greater) Len() int           { return len(g) }
func (g greater) Less(i, j int) bool { return g[i] > g[j] }
func (g greater) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }

func (con *contest) main() error {
	var N, A, B int
	con.Scan(&N, &A, &B)
	h := make([]int, N)
	hi := 0
	lo := 1
	for i := 0; i < N; i++ {
		con.Scan(&h[i])
		hi += (h[i] + A - 1) / A
	}
	sort.Sort(greater(h))
	for hi > lo {
		mi := (hi + lo) / 2
		a := 0
		for _, x := range h {
			// A*a + B*b > x
			// a + b = mi
			// (A - B) * a > x - B * mi
			x -= B * mi
			if x < 0 {
				break
			}
			x += A - B - 1
			x /= A - B
			a += x
		}
		if a > mi {
			if lo == mi {
				lo++
			} else {
				lo = mi
			}
		} else {
			hi = mi
		}
	}
	con.Println(lo)
	return nil
}
