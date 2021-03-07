package main

// Solution for https://atcoder.jp/contests/abc194/tasks/abc194_e

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
	var N, M int
	con.Scan(&N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&A[i])
	}
	S := make([]int, N+2)
	for i := 0; i < M; i++ {
		S[A[i]]++
	}
	var X, minX int
	for S[X] != 0 {
		X++
	}
	minX = X
	for i := M; i < N; i++ {
		a := A[i-M]
		b := A[i]
		S[a]--
		S[b]++
		if a < X && S[a] == 0 {
			X = a
			if X < minX {
				minX = X
			}
		}
		for S[X] != 0 {
			X++
		}
	}
	con.Println(minX)
	return nil
}
