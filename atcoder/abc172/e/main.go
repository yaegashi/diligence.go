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

const mod = 1e9 + 7

func (con *contest) main() error {
	var n, m int
	con.Scan(&n, &m)
	s := make([]int, 500001)
	k := m - n
	s[1] = k
	s[2] = k + 1 + k*k
	s[2] %= mod
	for i := 3; i <= n; i++ {
		x := s[i-2] * (i - 1)
		x %= mod
		y := s[i-1] * (i - 1 + k)
		y %= mod
		s[i] = x + y
		s[i] %= mod
	}
	p := 1
	for i := m; i > m-n; i-- {
		p *= i
		p %= mod
	}
	p *= s[n]
	p %= mod
	con.Println(p)
	return nil
}
