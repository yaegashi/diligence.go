package main

// Solution for https://atcoder.jp/contests/abc175/tasks/abc175_c

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
	var x, k, d int
	con.Scan(&x, &k, &d)
	if x < 0 {
		x = -x
	}
	y := x / d
	if y > k {
		y = k
	}
	x -= y * d
	k -= y
	if k == 0 {
		con.Println(x)
		return nil
	}
	if x == 0 {
		if k%2 == 0 {
			con.Println(0)
		} else {
			con.Println(d)
		}
		return nil
	}
	if k%2 == 0 {
		con.Println(x)
	} else {
		con.Println(d - x)
	}
	return nil
}
