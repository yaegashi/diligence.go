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
	var n int
	con.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&x[i], &y[i])
	}
	minCost := make([]int, n)
	used := make([]bool, n)
	res := 0
	for i := 1; i < n; i++ {
		minCost[i] = 1e18
	}
	for {
		v := -1
		for u := 0; u < n; u++ {
			if used[u] {
				continue
			}
			if v == -1 || minCost[u] < minCost[v] {
				v = u
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		res += minCost[v]
		for u := 0; u < n; u++ {
			cx := x[u] - x[v]
			if cx < 0 {
				cx = -cx
			}
			cy := y[u] - y[v]
			if cy < 0 {
				cy = -cy
			}
			if cx > cy {
				cx = cy
			}
			if cx < minCost[u] {
				minCost[u] = cx
			}
		}
	}
	con.Println(res)
	return nil
}
