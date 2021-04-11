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
	bRange := make([]ballRange, 0, N)
	for _, p := range bPos {
		n := len(bRange) - 1
		if n >= 0 && bRange[n].c == p.c {
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
	dTotal := int(1e18)
	var dfs func(x, n, l int)
	dfs = func(x, n, d int) {
		if d >= dTotal {
			return
		}
		if n == len(bRange) {
			if x > 0 {
				d += x
			} else {
				d -= x
			}
			if d < dTotal {
				dTotal = d
			}
			return
		}
		d += bRange[n].max - bRange[n].min
		dMax := bRange[n].max - x
		if dMax < 0 {
			dMax = -dMax
		}
		dfs(bRange[n].min, n+1, d+dMax)
		dMin := bRange[n].min - x
		if dMin < 0 {
			dMin = -dMin
		}
		dfs(bRange[n].max, n+1, d+dMin)
	}
	dfs(0, 0, 0)
	con.Println(dTotal)
	return nil
}

type ballPos struct{ x, c int }
type ballRange struct{ c, min, max int }

const maxPos = 1e10
const minPos = -1e10
