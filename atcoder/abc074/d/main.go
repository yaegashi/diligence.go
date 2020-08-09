package main

// Solution for https://atcoder.jp/contests/abc074/tasks/arc083_b

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

const INF = 1 << 50

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	A := make([][]int, N)
	B := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		B[i] = make([]int, N)
		for j := 0; j < N; j++ {
			con.Scan(&A[i][j])
			B[i][j] = A[i][j]
		}
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				x := A[i][k] + A[k][j]
				if x < A[i][j] {
					A[i][j] = x
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if A[i][j] != B[i][j] {
				con.Println(-1)
				return nil
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := 0; k < N; k++ {
				if k == i || k == j {
					continue
				}
				if A[i][k]+A[k][j] == A[i][j] {
					B[i][j] = INF
					B[j][i] = INF
					break
				}
			}
		}
	}
	L := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if B[i][j] != INF {
				L += B[i][j]
			}
		}
	}
	con.Println(L)
	return nil
}
