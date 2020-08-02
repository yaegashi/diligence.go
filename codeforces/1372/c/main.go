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
	for testcase := 0; testcase < T; testcase++ {
		var n int
		con.Scan(&n)
		x := make([]bool, n)
		c := 0
		for i := 0; i < n; i++ {
			var a int
			con.Scan(&a)
			if a == i+1 {
				c++
				x[i] = true
			}
		}
		if c == n {
			con.Println(0)
			continue
		}
		p := 0
		for i := 0; i < n; i++ {
			if !x[i] && (i == 0 || x[i-1]) {
				p++
			}
		}
		if p == 1 {
			con.Println(1)
		} else {
			con.Println(2)
		}
	}
	return nil
}
