package main

// Solution for https://atcoder.jp/contests/abc173/tasks/abc173_d

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

func (con *contest) main() error {
	var n int
	con.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	b := make([]int, 0, n*2)
	b = append(b, a[0])
	c := 0
	for i := 1; i < n; i++ {
		c += a[i-1]
		b = append(b, a[i])
		b = append(b, a[i])
	}
	con.Println(c)
	return nil
}
