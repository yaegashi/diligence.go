package main

// Solution for https://atcoder.jp/contests/abc195/tasks/abc195_d

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

type pt struct{ w, v, u int }

func (con *contest) main() error {
	var N, M, Q int
	con.Scan(&N, &M, &Q)
	P := make([]pt, N)
	X := make([]int, M)
	for i := 0; i < N; i++ {
		con.Scan(&P[i].w, &P[i].v)
	}
	for i := 0; i < M; i++ {
		con.Scan(&X[i])
	}
	sort.Slice(P, func(i, j int) bool { return P[i].v > P[j].v })
	for i := 0; i < Q; i++ {
		var L, R int
		con.Scan(&L, &R)
		B := make([]int, 0, len(X))
		for j := 0; j < len(X); j++ {
			if j < L-1 || j > R-1 {
				B = append(B, X[j])
			}
		}
		sort.Ints(B)
		T := 0
		for p := 0; p < len(P); p++ {
			P[p].u = 0
		}
		for _, b := range B {
			for p := 0; p < len(P); p++ {
				if P[p].u == 0 && P[p].w <= b {
					P[p].u = 1
					T += P[p].v
					break
				}
			}
		}
		con.Println(T)
	}
	return nil
}
