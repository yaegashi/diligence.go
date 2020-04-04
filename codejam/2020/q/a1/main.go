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

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var N int
		con.Scan(&N)
		M := make([][]int, N)
		for i := 0; i < N; i++ {
			M[i] = make([]int, N)
			for j := 0; j < N; j++ {
				con.Scan(&M[i][j])
			}
		}
		k := 0
		r := 0
		c := 0
		for i := 0; i < N; i++ {
			k += M[i][i]
			R := make([]bool, N)
			for j := 0; j < N; j++ {
				if R[M[i][j]-1] {
					r++
					break
				}
				R[M[i][j]-1] = true
			}
			C := make([]bool, N)
			for j := 0; j < N; j++ {
				if C[M[j][i]-1] {
					c++
					break
				}
				C[M[j][i]-1] = true
			}
		}
		con.Printf("Case #%d: %d %d %d\n", testcase, k, r, c)
	}
	return nil
}
