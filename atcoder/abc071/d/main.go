package main

// Solution for https://atcoder.jp/contests/abc071/tasks/arc081_b

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

const mod = 1000000007

func (con *contest) main() error {
	var n int
	var s1, s2 string
	con.Scan(&n, &s1, &s2)
	p := 1
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			if i == 0 {
				p *= 3
				p %= mod
			} else if s1[i-1] == s2[i-1] {
				p *= 2
				p %= mod
			}
		} else {
			if i == 0 {
				p *= 6
				p %= mod
			} else if s1[i-1] == s2[i-1] {
				p *= 2
				p %= mod
			} else {
				p *= 3
				p %= mod
			}
			i++
		}
	}
	con.Println(p)
	return nil
}
