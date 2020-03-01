package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type contest struct {
	in      io.Reader
	out     io.Writer
	N, M, A int
	D       []bool
	E       [][]bool
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
	con.Scan(&con.N, &con.M)
	con.D = make([]bool, con.N)
	con.E = make([][]bool, con.N)
	for i := range con.E {
		con.E[i] = make([]bool, con.N)
	}
	for i := 0; i < con.M; i++ {
		var a, b int
		con.Scan(&a, &b)
		con.E[a-1][b-1] = true
		con.E[b-1][a-1] = true
	}
	con.rec(0, 0)
	con.Println(con.A)
	return nil
}

func (con *contest) rec(v, n int) {
	if n+1 == con.N {
		con.A++
		return
	}
	con.D[v] = true
	for i := 0; i < con.N; i++ {
		if con.E[v][i] && !con.D[i] {
			con.rec(i, n+1)
		}
	}
	con.D[v] = false
}
