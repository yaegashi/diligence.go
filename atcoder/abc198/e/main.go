package main

// Solution for https://atcoder.jp/contests/abc198/tasks/abc198_e

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

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	C := make([]int, N)
	E := make([][]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&C[i])
	}
	for i := 1; i < N; i++ {
		var a, b int
		con.Scan(&a, &b)
		a--
		b--
		E[a] = append(E[a], b)
		E[b] = append(E[b], a)
	}
	F := make([][]int, N)
	V := make([]bool, N)
	Q := make([]int, 0, N*3)
	Q = append(Q, 0)
	for len(Q) > 0 {
		x := Q[0]
		Q = Q[1:]
		V[x] = true
		for _, y := range E[x] {
			if V[y] {
				continue
			}
			F[x] = append(F[x], y)
			Q = append(Q, y)
		}
	}
	M := make([]int, 0, N)
	CM := map[int]bool{}
	var rec func(x int)
	rec = func(x int) {
		c := C[x]
		if CM[c] {
			for _, y := range F[x] {
				rec(y)
			}
		} else {
			M = append(M, x+1)
			CM[c] = true
			for _, y := range F[x] {
				rec(y)
			}
			CM[c] = false
		}
	}
	rec(0)
	sort.Ints(M)
	for _, m := range M {
		con.Println(m)
	}
	return nil
}
