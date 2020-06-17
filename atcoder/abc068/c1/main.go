package main

// Solution for https://atcoder.jp/contests/abc068/tasks/arc079_a

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
	var n, m int
	con.Scan(&n, &m)
	x := make([]bool, n+1)
	y := make([]bool, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		var z bool
		con.Scan(&a, &b)
		if a == 1 {
			x[b] = true
			z = y[b]
		}
		if a == n {
			y[b] = true
			z = x[b]
		}
		if b == 1 {
			x[a] = true
			z = y[a]
		}
		if b == n {
			y[a] = true
			z = x[a]
		}
		if z {
			con.Println("POSSIBLE")
			return nil
		}
	}
	con.Println("IMPOSSIBLE")
	return nil
}
