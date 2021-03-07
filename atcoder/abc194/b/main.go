package main

// Solution for https://atcoder.jp/contests/abc194/tasks/abc194_b

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
	var N int
	con.Scan(&N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&A[i], &B[i])
	}
	min := 200000
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			var x int
			if i == j {
				x = A[i] + B[j]
			} else {
				x = A[i]
				if B[j] > x {
					x = B[j]
				}
			}
			if x < min {
				min = x
			}
		}
	}
	con.Println(min)
	return nil
}
