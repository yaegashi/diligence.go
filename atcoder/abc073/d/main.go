package main

// Solution for https://atcoder.jp/contests/abc073/tasks/abc073_d

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

const INF = 1 << 32

func (con *contest) main() error {
	var N, M, R int
	con.Scan(&N, &M, &R)
	r := make([]int, R)
	for i := 0; i < R; i++ {
		con.Scan(&r[i])
		r[i]--
	}
	E := make([][]int, N)
	for i := 0; i < N; i++ {
		E[i] = make([]int, N)
		for j := 0; j < N; j++ {
			E[i][j] = INF
		}
	}
	for i := 0; i < N; i++ {
		E[i][i] = 0
	}
	for i := 0; i < M; i++ {
		var a, b, c int
		con.Scan(&a, &b, &c)
		E[a-1][b-1] = c
		E[b-1][a-1] = c
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				x := E[i][k] + E[k][j]
				if x < E[i][j] {
					E[i][j] = x
				}
			}
		}
	}
	P := make([][]int, R)
	for i := 0; i < R; i++ {
		P[i] = make([]int, 1<<R)
		for j := 0; j < 1<<R; j++ {
			P[i][j] = INF
		}
	}
	var dfs func(int, int)
	dfs = func(p int, x int) {
		for q := 0; q < R; q++ {
			y := x | (1 << q)
			if x == y {
				continue
			}
			a := E[r[p]][r[q]]
			if P[p][x]+a < P[q][y] {
				P[q][y] = P[p][x] + a
				dfs(q, y)
			}
		}
	}
	for i := 0; i < R; i++ {
		P[i][1<<i] = 0
		dfs(i, 1<<i)
	}
	min := INF
	for i := 0; i < R; i++ {
		if P[i][(1<<R)-1] < min {
			min = P[i][(1<<R)-1]
		}
	}
	con.Println(min)
	return nil
}
