package main

// Solution for https://atcoder.jp/contests/abc190/tasks/abc190_c

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
	var N, M, K int
	con.Scan(&N, &M)
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		con.Scan(&A[i], &B[i])
		A[i]--
		B[i]--
	}
	con.Scan(&K)
	C := make([]int, K)
	D := make([]int, K)
	for i := 0; i < K; i++ {
		con.Scan(&C[i], &D[i])
		C[i]--
		D[i]--
	}
	maxS := 0
	for i := 0; i < 1<<K; i++ {
		S := 0
		X := make([]bool, N)
		for j := 0; j < K; j++ {
			if i&(1<<j) == 0 {
				X[C[j]] = true
			} else {
				X[D[j]] = true
			}
		}
		for j := 0; j < M; j++ {
			if X[A[j]] && X[B[j]] {
				S++
			}
		}
		if S > maxS {
			maxS = S
		}
	}
	con.Println(maxS)
	return nil
}
