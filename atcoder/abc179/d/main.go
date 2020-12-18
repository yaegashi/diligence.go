package main

// Solution for https://atcoder.jp/contests/abc179/tasks/abc179_d

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

const M = 998244353

func (con *contest) main() error {
	var N, K int
	con.Scan(&N, &K)
	s := make([]int, 0, 200000)
	for i := 0; i < K; i++ {
		var l, r int
		con.Scan(&l, &r)
		for j := l; j <= r; j++ {
			s = append(s, j)
		}
	}
	sort.Ints(s)
	x := make([]int, N)
	x[0] = 1
	for i := 1; i < N; i++ {
		for _, a := range s {
			if a > i {
				break
			}
			x[i] += x[i-a]
			x[i] %= M
		}
	}
	con.Println(x[N-1])
	return nil
}
