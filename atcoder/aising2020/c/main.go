package main

// Solution for https://atcoder.jp/contests/aising2020/tasks/aising2020_c

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
	var n int
	con.Scan(&n)
	a := make([]int, n+1)
	for x := 1; x <= n; x++ {
		p := x * x
		if p > n {
			break
		}
		for y := 1; y <= n; y++ {
			q := p + y*y + x*y
			if q > n {
				break
			}
			for z := 1; z <= n; z++ {
				r := q + z*z + y*z + z*x
				if r > n {
					break
				}
				a[r]++
			}
		}
	}
	for i := 1; i <= n; i++ {
		con.Println(a[i])
	}
	return nil
}
