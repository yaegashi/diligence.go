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
	for i := 0; i <= N; i++ {
		D[i] = make([]int, 100000)
	}
	D[0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; D[i][j] > 0; j++ {
			for k := 0; k <= K; k++ {
				D[i+1][(i+1)*k+j] += D[i][j]
			}
		}
	}
	//	for i := 0; i < N; i++ {
	//		for j := 1; D[i][j] > 0; j++ {
	//			D[i+1][j] += D[i][j]
	//		}
	//	}
	X := make([]int, N)
	X[0] = K
	X[N-1] = K
	for i := 1; i < N-1; i++ {
		X[i] = K
		a := i
		b := N - i - 1
		for j := 1; D[a][j] > 0 && D[b][j] > 0; j++ {
			x := K + 1
			x *= D[a][j]
			x %= M
			x *= D[b][j]
			x %= M
			X[i] += x
			X[i] %= M
		}
	}
	for i := 0; i < N; i++ {
		con.Println(X[i])
	}
	return nil
}
