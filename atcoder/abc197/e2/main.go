package main

// Solution for https://atcoder.jp/contests/abc197/tasks/abc197_e

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
	var N int
	con.Scan(&N)
	bPos := make([]ballPos, N)
	for i := 0; i < N; i++ {
		con.Scan(&bPos[i].x, &bPos[i].c)
	}
	sort.Slice(bPos, func(i, j int) bool { return bPos[i].c < bPos[j].c })
	bRange := []ballRange{{0, 0, 0}}
	for _, p := range bPos {
		n := len(bRange) - 1
		if bRange[n].c == p.c {
			if p.x < bRange[n].min {
				bRange[n].min = p.x
			}
			if p.x > bRange[n].max {
				bRange[n].max = p.x
			}
		} else {
			bRange = append(bRange, ballRange{p.c, p.x, p.x})
		}
	}
	bRange = append(bRange, ballRange{0, 0, 0})
	bMemo := map[ballKey]int{}
	var dfs func(m bool, n int) int
	dfs = func(m bool, n int) int {
		if n == len(bRange)-1 {
			return 0
		}
		d, ok := bMemo[ballKey{m, n}]
		if !ok {
			d0 := dfs(false, n+1)
			d1 := dfs(true, n+1)
			var x0, x1 int
			if m {
				x0 = bRange[n+1].min - bRange[n].min
				if x0 < 0 {
					x0 = -x0
				}
				x0 += d0
				x1 = bRange[n+1].max - bRange[n].min
				if x1 < 0 {
					x1 = -x1
				}
				x1 += d1
			} else {
				x0 = bRange[n+1].min - bRange[n].max
				if x0 < 0 {
					x0 = -x0
				}
				x0 += d0
				x1 = bRange[n+1].max - bRange[n].max
				if x1 < 0 {
					x1 = -x1
				}
				x1 += d1
			}
			d = bRange[n].max - bRange[n].min
			if x0 < x1 {
				d += x0
			} else {
				d += x1
			}
			bMemo[ballKey{m, n}] = d
		}
		return d
	}
	con.Println(dfs(false, 0))
	return nil
}

type ballPos struct{ x, c int }
type ballRange struct{ c, min, max int }
type ballKey struct {
	m bool
	n int
}

const maxPos = 1e10
const minPos = -1e10
