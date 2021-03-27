package main

// Solution for https://atcoder.jp/contests/abc196/tasks/abc196_e

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

const MAX = 10000000000
const MIN = -10000000000

type OP struct{ add, min, max int }

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	A := make([]int, N)
	T := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&A[i], &T[i])
	}
	var Q int
	con.Scan(&Q)
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		con.Scan(&X[i])
	}
	O := make([]OP, 0, N)
	o := OP{add: 0, min: MIN, max: MAX}
	for i := 0; i < N; i++ {
		switch T[i] {
		case 1:
			if o.min != MIN || o.max != MAX {
				O = append(O, o)
				o = OP{add: 0, min: MIN, max: MAX}
			}
			o.add += A[i]
		case 2:
			if A[i] > o.min {
				o.min = A[i]
			}
		case 3:
			if A[i] < o.max {
				o.max = A[i]
			}
		}
	}
	O = append(O, o)
	for _, x := range X {
		for i := 0; i < len(O); i++ {
			x += O[i].add
			if x > O[i].max {
				x = O[i].max
			}
			if x < O[i].min {
				x = O[i].min
			}
		}
		con.Println(x)
	}
	return nil
}
