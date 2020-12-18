package main

// Solution for https://atcoder.jp/contests/abc183/tasks/abc183_c

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
	T := make([][]int, N)
	for i := 0; i < N; i++ {
		T[i] = make([]int, N)
		for j := 0; j < N; j++ {
			con.Scan(&T[i][j])
		}
	}
	C := 0
	B := make([]bool, N)
	var dfs func(int, int)
	dfs = func(n int, k int) {
		B[n] = true
		done := true
		for i := 0; i < N; i++ {
			if B[i] {
				continue
			}
			done = false
			dfs(i, k+T[n][i])
		}
		if done && k+T[n][0] == K {
			C++
		}
		B[n] = false
	}
	dfs(0, 0)
	con.Println(C)
	return nil
}
