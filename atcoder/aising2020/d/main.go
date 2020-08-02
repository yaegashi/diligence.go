package main

// Solution for https://atcoder.jp/contests/aising2020/tasks/aising2020_d

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

func f(x int) int {
	if x == 0 {
		return 0
	}
	p := x
	p = ((p & 0xaaaaaaaa) >> 1) + (p & 0x55555555)
	p = ((p & 0xcccccccc) >> 2) + (p & 0x33333333)
	p = ((p & 0xf0f0f0f0) >> 4) + (p & 0x0f0f0f0f)
	p = ((p & 0xff00ff00) >> 8) + (p & 0x00ff00ff)
	p = ((p & 0xffff0000) >> 16) + (p & 0x0000ffff)
	return 1 + f(x%p)
}

func (con *contest) main() error {
	var N int
	var X string
	con.Scan(&N, &X)
	C := 0
	for i := 0; i < N; i++ {
		if X[i] == '1' {
			C++
		}
	}
	b1 := make([]int, N)
	b2 := make([]int, N)
	b1[0] = 1 % (C + 1)
	if C > 1 {
		b2[0] = 1 % (C - 1)
	}
	for i := 1; i < N; i++ {
		b1[i] = b1[i-1] * 2
		b1[i] %= C + 1
		if C > 1 {
			b2[i] = b2[i-1] * 2
			b2[i] %= C - 1
		}
	}
	x1, x2 := 0, 0
	for i := 0; i < N; i++ {
		if X[N-1-i] == '1' {
			x1 += b1[i]
			x1 %= C + 1
			if C > 1 {
				x2 += b2[i]
				x2 %= C - 1
			}
		}
	}
	for i := 0; i < N; i++ {
		z := 0
		if X[i] == '0' {
			z = x1 + b1[N-1-i]
			z %= C + 1
			z = 1 + f(z)
		} else if C > 1 {
			z = x2 - b2[N-1-i]
			if z < 0 {
				z += C - 1
			}
			z = 1 + f(z)
		} else {
			z = 0
		}
		con.Println(z)
	}
	return nil
}
