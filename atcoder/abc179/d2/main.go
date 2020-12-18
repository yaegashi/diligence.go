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

type seg struct{ l, r int }

func (con *contest) main() error {
	var N, K int
	con.Scan(&N, &K)
	segs := make([]seg, K)
	for i := 0; i < K; i++ {
		con.Scan(&segs[i].l, &segs[i].r)
		segs[i].r++
	}
	sort.Slice(segs, func(i, j int) bool { return segs[i].l < segs[j].r })
	x := make([]int, N)
	y := make([]int, N)
	x[0] = 1
	y[0] = 1
	for i := 1; i < N; i++ {
		z := 0
		for _, s := range segs {
			if s.l > i {
				break
			}
			z += y[i-s.l]
			z %= M
			if s.r <= i {
				z -= y[i-s.r]
				if z < 0 {
					z += M
				}
			}
		}
		x[i] = z
		y[i] = y[i-1] + x[i]
		y[i] %= M
	}
	con.Println(x[N-1])
	return nil
}
