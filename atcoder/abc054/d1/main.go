package main

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

const eMax = 200000
const eOffset = 100000

func (con *contest) main() error {
	var N, Ma, Mb int
	con.Scan(&N, &Ma, &Mb)
	C := make([]int, N)
	D := make([]int, N)
	E := make([][]int, N)
	for i := 0; i < N; i++ {
		var a, b int
		con.Scan(&a, &b, &C[i])
		D[i] = a*Mb - b*Ma
	}
	E[0] = make([]int, eMax)
	E[0][D[0]+eOffset] = C[0]
	for i := 1; i < N; i++ {
		E[i] = make([]int, eMax)
		for j := 0; j < eMax; j++ {
			k := j - D[i]
			if 0 <= k && k < eMax && E[i-1][k] > 0 && (E[i-1][j] == 0 || E[i-1][k]+C[i] < E[i-1][j]) {
				E[i][j] = E[i-1][k] + C[i]
			} else {
				E[i][j] = E[i-1][j]
			}
		}
		if E[i][D[i]+eOffset] == 0 || C[i] < E[i][D[i]+eOffset] {
			E[i][D[i]+eOffset] = C[i]
		}
	}
	A := -1
	for i := 1; i < N; i++ {
		if E[i][eOffset] > 0 && (A < 0 || E[i][eOffset] < A) {
			A = E[i][eOffset]
		}
	}
	con.Println(A)
	return nil
}
