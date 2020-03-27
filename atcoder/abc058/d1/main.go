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

func (con *contest) sum(x []int) int {
	n := len(x) - 1
	p := n
	q := 0
	r := 0
	for i := 0; p > 0; i++ {
		q += p
		s := ((x[i+1] - x[i]) * q) % mod
		r = (r + s) % mod
		if p > 1 {
			t := ((x[n-i] - x[n-i-1]) * q) % mod
			r = (r + t) % mod
		}
		p -= 2
	}
	return r
}

func (con *contest) main() error {
	var n, m int
	con.Scan(&n, &m)
	x := make([]int, n)
	y := make([]int, m)
	for i := 0; i < n; i++ {
		con.Scan(&x[i])
	}
	for i := 0; i < m; i++ {
		con.Scan(&y[i])
	}
	sumX := con.sum(x)
	sumY := con.sum(y)
	con.Println((sumX * sumY) % mod)
	return nil
}

/*
012345678
111111111 9
122222221 16
123333321 21
123444321 24
123454321 25
123444321 24
123333321 21
122222221 16
111111111 9

01234567
11111111 8
12222221 14
12333321 18
12344321 20
12344321 20
12333321 18
12222221 14
11111111 8
*/
