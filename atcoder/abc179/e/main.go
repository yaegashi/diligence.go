package main

// Solution for https://atcoder.jp/contests/abc179/tasks/abc179_e

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
	var N, X, M int
	con.Scan(&N, &X, &M)
	mods := make([]int, M)
	for i := 0; i < M; i++ {
		mods[i] = -1
	}
	var p, q, r, s int
	for r = X; mods[r] < 0; r = (r * r) % M {
		mods[r] = q
		q++
	}
	aTotal := 0
	aSum := make([]int, M+1)
	for s = X; s != r; s = (s * s) % M {
		aTotal += s
		aSum[p+1] = aSum[p] + s
		p++
	}
	bTotal := 0
	bSum := make([]int, M+1)
	for i := p; i < q; i++ {
		bTotal += s
		bSum[i-p+1] = bSum[i-p] + s
		s = (s * s) % M
	}
	if N < p {
		con.Println(aSum[N])
	} else {
		x := (N - p) / (q - p)
		y := (N - p) % (q - p)
		con.Println(aTotal + bTotal*x + bSum[y])
	}
	return nil
}
