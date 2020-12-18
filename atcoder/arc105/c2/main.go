package main

// Solution for https://atcoder.jp/contests/arc105/tasks/arc105_c

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

type bridge struct {
	l, v int
}

func rec(v int, t []int, w []int) bool {
	if len(w) == 0 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i]+w[0] <= v {
			t[i] += w[0]
			if rec(v, t, w[1:]) {
				return true
			}
			t[i] -= w[0]
		}
	}
	return false
}

func (con *contest) main() error {
	var N, M int
	con.Scan(&N, &M)
	w := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&w[i])
	}
	b := make([]bridge, M)
	for i := 0; i < M; i++ {
		con.Scan(&b[i].l, &b[i].v)
	}
	sort.Slice(b, func(i, j int) bool {
		if b[i].v == b[j].v {
			return b[i].l > b[j].l
		} else {
			return b[i].v < b[j].v
		}
	})
	maxS := 0
	minS := 1
	for i := 0; i < M; i++ {
		ok := false
		for s := minS; s <= N; s++ {
			t := make([]int, s)
			if rec(b[i].v, t, w) {
				S := b[i].l * (s - 1)
				if S > maxS {
					maxS = S
					minS = s
				}
				ok = true
				break
			}
		}
		if !ok {
			con.Println(-1)
			return nil
		}
	}
	con.Println(maxS)
	return nil
}
