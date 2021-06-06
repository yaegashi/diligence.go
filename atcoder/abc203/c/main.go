package main

// Solution for https://atcoder.jp/contests/abc203/tasks/abc203_c

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

type qab struct{ a, b int }

func (con *contest) main() error {
	var n, k int
	con.Scan(&n, &k)
	q := make([]qab, n)
	for i := 0; i < n; i++ {
		con.Scan(&q[i].a, &q[i].b)
	}
	sort.Slice(q, func(i, j int) bool { return q[i].a < q[j].a })
	p := 0
	for i := 0; i < n; i++ {
		a, b := q[i].a, q[i].b
		if a-p > k {
			break
		}
		k -= a - p
		p = a
		k += b
	}
	p += k
	con.Println(p)
	return nil
}
