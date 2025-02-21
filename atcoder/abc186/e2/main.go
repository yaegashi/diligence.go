package main

// Solution for https://atcoder.jp/contests/abc186/tasks/abc186_e

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

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func inverse(a, m int) int {
	b := m
	x, y := 1, 0
	for b > 0 {
		q := a / b
		a, b = b, a-q*b
		x, y = y, x-q*y
	}
	if x < 0 {
		x += m
	}
	return x
}

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for t := 0; t < T; t++ {
		var N, S, K int
		con.Scan(&N, &S, &K)
		K %= N
		d := gcd(N, S)
		d = gcd(d, K)
		N /= d
		S /= d
		K /= d
		if gcd(N, K) != 1 {
			con.Println(-1)
			continue
		}
		z := (N - S) * inverse(K, N)
		z %= N
		con.Println(z)
	}
	return nil
}
