package main

// Solution for https://atcoder.jp/contests/abc197/tasks/abc197_c

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
	for i := 0; i < N; i++ {
		con.Scan(&A[i])
	}
	minX := 1 << 31
	for i := 0; i < 1<<N; i++ {
		X := calc(A, i)
		if X < minX {
			minX = X
		}
	}
	con.Println(minX)
	return nil
}

func calc(A []int, B int) int {
	X, Y := 0, 0
	for i := 0; i < len(A); i++ {
		X |= A[i]
		if B&(1<<i) != 0 {
			Y ^= X
			X = 0
		}
	}
	return Y ^ X
}
