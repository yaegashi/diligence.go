package main

// Solution for https://atcoder.jp/contests/abc174/tasks/abc174_e

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
	var n, k int
	con.Scan(&n, &k)
	a := make([]int, n)
	lo := 1
	hi := 0
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
		if a[i] > hi {
			hi = a[i]
		}
	}
	for lo < hi {
		mid := (lo + hi) / 2
		c := 0
		for i := 0; i < n; i++ {
			c += (a[i] - 1) / mid
		}
		if c <= k {
			hi = mid
		} else if lo == mid {
			lo++
		} else {
			lo = mid
		}
	}
	con.Println((lo + hi) / 2)
	return nil
}
