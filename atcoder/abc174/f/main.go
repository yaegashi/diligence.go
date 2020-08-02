package main

// Solution for https://atcoder.jp/contests/abc174/tasks/abc174_f

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

type query struct{ l, r, i int }

func update(i int, v int, bit []int) {
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func ask(i int, bit []int) int {
	s := 0
	for i > 0 {
		s += bit[i]
		i -= i & -i
	}
	return s
}

func (con *contest) main() error {
	var n, m int
	con.Scan(&n, &m)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&c[i])
		c[i]--
	}
	q := make([]query, m)
	for i := 0; i < m; i++ {
		con.Scan(&q[i].l, &q[i].r)
		q[i].l--
		q[i].r--
		q[i].i = i
	}
	sort.Slice(q, func(i, j int) bool { return q[i].r < q[j].r })
	bit := make([]int, n+1)
	last := make([]int, n+1)
	for i := 0; i <= n; i++ {
		last[i] = -1
	}
	ans := make([]int, m)
	j := 0
	for i := 0; i < n; i++ {
		if last[c[i]] != -1 {
			update(last[c[i]]+1, -1, bit)
		}
		last[c[i]] = i
		update(i+1, 1, bit)
		for j < m && q[j].r == i {
			ans[q[j].i] = ask(q[j].r+1, bit) - ask(q[j].l, bit)
			j++
		}
	}
	for i := 0; i < m; i++ {
		con.Println(ans[i])
	}
	return nil
}
