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

func (con *contest) main() error {
	var N, K, L int
	con.Scan(&N, &K, &L)
	roads := make([]int, N)
	rails := make([]int, N)
	for i := range roads {
		roads[i] = i
		rails[i] = i
	}
	for i := 0; i < K; i++ {
		var p, q int
		con.Scan(&p, &q)
		p--
		q--
		for p != roads[p] {
			p = roads[p]
		}
		for q != roads[q] {
			q = roads[q]
		}
		roads[p] = roads[q]
	}
	for i := 0; i < L; i++ {
		var p, q int
		con.Scan(&p, &q)
		p--
		q--
		for p != rails[p] {
			p = rails[p]
		}
		for q != rails[q] {
			q = rails[q]
		}
		rails[p] = rails[q]
	}
	cmap := map[struct{ a, b int }]int{}
	for i := 0; i < N; i++ {
		for roads[i] != roads[roads[i]] {
			roads[i] = roads[roads[i]]
		}
		for rails[i] != rails[rails[i]] {
			rails[i] = rails[rails[i]]
		}
		cmap[struct{ a, b int }{roads[i], rails[i]}]++
	}
	a := make([]interface{}, N)
	for i := 0; i < N; i++ {
		a[i] = cmap[struct{ a, b int }{roads[i], rails[i]}]
	}
	con.Println(a...)
	return nil
}
