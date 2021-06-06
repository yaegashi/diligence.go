package main

// Solution for https://atcoder.jp/contests/abc203/tasks/abc203_d

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
	var N, K int
	con.Scan(&N, &K)
	M := make([][]int, N)
	for i := 0; i < N; i++ {
		M[i] = make([]int, N)
		for j := 0; j < N; j++ {
			con.Scan(&M[i][j])
		}
	}
	S := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		S[i] = make([]int, N+1)
	}
	Scalc := func(X int) {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				S[i][j] = S[i-1][j] + S[i][j-1] - S[i-1][j-1]
				if M[i-1][j-1] > X {
					S[i][j]++
				}
			}
		}
	}
	lo := -1
	hi := int(1e9 + 1)
	for lo+1 < hi {
		mi := (lo + hi) / 2
		Scalc(mi)
		ok := false
		for i := 0; i <= N-K; i++ {
			for j := 0; j <= N-K; j++ {
				s := S[i+K][j+K] + S[i][j] - S[i+K][j] - S[i][j+K]
				if s <= K*K/2 {
					ok = true
					break
				}
			}
		}
		if ok {
			hi = mi
		} else {
			lo = mi
		}
	}
	con.Println(hi)
	return nil
}
