package main

// Solution for https://atcoder.jp/contests/abc193/tasks/abc193_f

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

func score(M [][]int, z int) int {
	N := len(M)
	S := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a := M[i][j]
			if a < 0 {
				continue
			}
			if a == 2 {
				if (i+j)&1 == z {
					a = 0
				} else {
					a = 1
				}
			}
			if i < N-1 {
				b := M[i+1][j]
				if b >= 0 && a != b {
					S++
				}
			}
			if j < N-1 {
				c := M[i][j+1]
				if c >= 0 && a != c {
					S++
				}
			}
		}
	}
	return S
}

func fill2(M [][]int, p int, q int) {
	N := len(M)
	M[p][q] = 2
	if p > 0 && M[p-1][q] < 0 {
		fill2(M, p-1, q)
	}
	if q > 0 && M[p][q-1] < 0 {
		fill2(M, p, q-1)
	}
	if p < N-1 && M[p+1][q] < 0 {
		fill2(M, p+1, q)
	}
	if q < N-1 && M[p][q+1] < 0 {
		fill2(M, p, q+1)
	}
}

func fillz(M [][]int, z int) {
	N := len(M)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if M[i][j] == 2 {
				if (i+j)&1 == z {
					M[i][j] = 0
				} else {
					M[i][j] = 1
				}
			}
		}
	}
}

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	M := make([][]int, N)
	for i := 0; i < N; i++ {
		var S string
		con.Scan(&S)
		M[i] = make([]int, N)
		for j, a := range S {
			switch a {
			case 'B':
				M[i][j] = 0
			case 'W':
				M[i][j] = 1
			case '?':
				M[i][j] = -1
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if M[i][j] < 0 {
				fill2(M, i, j)
				if score(M, 0) > score(M, 1) {
					fillz(M, 0)
				} else {
					fillz(M, 1)
				}
			}
		}
	}
	con.Println(score(M, 0))
	return nil
}
