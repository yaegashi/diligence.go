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
	var testcases int
	con.Scan(&testcases)
	for testcase := 0; testcase < testcases; testcase++ {
		var n, m int
		con.Scan(&n, &m)
		r := make([]string, n)
		for i := 0; i < n; i++ {
			con.Scan(&r[i])
		}
		c := 0
		for i := 0; i < n-1; i++ {
			if r[i][m-1] != 'D' {
				c++
			}
		}
		for i := 0; i < m-1; i++ {
			if r[n-1][i] != 'R' {
				c++
			}
		}
		con.Println(c)
	}
	return nil
}
