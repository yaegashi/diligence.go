package main

// Solution for https://atcoder.jp/contests/abc183/tasks/abc183_f

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

type query struct{ o, a, b int }

func (con *contest) main() error {
	var N, Q int
	con.Scan(&N, &Q)
	C := make([]int, N)
	U := make([]int, N)
	M := make([]map[int]int, N)
	R := make([]query, Q)
	S := make([]bool, N)
	for i := 0; i < N; i++ {
		con.Scan(&C[i])
		C[i]--
		U[i] = i
	}
	for i := 0; i < Q; i++ {
		con.Scan(&R[i].o, &R[i].a, &R[i].b)
		R[i].a--
		R[i].b--
		if R[i].o == 2 {
			S[R[i].b] = true
		}
	}
	for _, r := range R {
		o, a, b := r.o, r.a, r.b
		if o == 1 {
			x := a
			for x != U[x] {
				x = U[x]
			}
			y := b
			for y != U[y] {
				y = U[y]
			}
			if x != y {
				mx := M[x]
				if mx == nil {
					mx = map[int]int{C[x]: 1}
				}
				my := M[y]
				if my == nil {
					my = map[int]int{C[y]: 1}
				}
				if len(mx) < len(my) {
					mx, my = my, mx
					x, y = y, x
				}
				for k, v := range my {
					if S[k] {
						mx[k] += v
					}
				}
				M[x] = mx
				M[y] = nil
			}
			for a != x {
				U[a], a = x, U[a]
			}
			for b != x {
				U[b], b = x, U[b]
			}
		} else {
			x := a
			for x != U[x] {
				x = U[x]
			}
			for a != x {
				U[a], a = x, U[a]
			}
			mx := M[x]
			if mx == nil {
				mx = map[int]int{C[x]: 1}
			}
			con.Println(mx[b])
		}
	}
	return nil
}
