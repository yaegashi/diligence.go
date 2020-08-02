package main

// Solution for https://atcoder.jp/contests/m-solutions2020/tasks/m_solutions2020_d

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
	A := make([]int, N+1)
	for i := 0; i < N; i++ {
		con.Scan(&A[i])
	}
	D := 0
	M := 1000
	S := 0
	if A[1] > A[0] {
		S += M / A[0]
		M %= A[0]
		D = 1
	}
	for i := 1; i < N; i++ {
		if D == 0 {
			if A[i+1] > A[i] {
				S += M / A[i]
				M %= A[i]
				D = 1
			}
		} else {
			if A[i+1] < A[i] {
				M += S * A[i]
				S = 0
				D = 0
			}
		}
	}
	con.Println(M)
	return nil
}
