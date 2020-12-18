package main

// Solution for https://atcoder.jp/contests/arc104/tasks/arc104_d

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
	var N, K, M int
	con.Scan(&N, &K, &M)
	D := make([][]int, N+1)
	L := K * N * (N + 1) / 2
	for i := 0; i <= N; i++ {
		D[i] = make([]int, L)
	}
	D[0][0] = 1
	for i := 0; i < N; i++ {
		J := K * i * (i + 1) / 2
		for j := 0; j <= J; j++ {
			for k := 0; k <= K; k++ {
				x := (i+1)*k + j
				if x >= L {
					break
				}
				D[i+1][x] += D[i][j]
				D[i+1][x] %= M
			}
		}
	}
	X := make([]int, N)
	X[0] = K
	X[N-1] = K
	for i := 1; i < N-1; i++ {
		a := i
		b := N - i - 1
		if a > b {
			break
		}
		X[i] = K
		J := K * a * (a + 1) / 2
		for j := 1; j <= J; j++ {
			x := K + 1
			x *= D[a][j]
			x %= M
			x *= D[b][j]
			x %= M
			X[i] += x
			X[i] %= M
		}
		X[N-i-1] = X[i]
	}
	for i := 0; i < N; i++ {
		con.Println(X[i])
	}
	return nil
}
