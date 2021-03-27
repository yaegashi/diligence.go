package main

// Solution for https://atcoder.jp/contests/abc196/tasks/abc196_d

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
	var H, W, A, B int
	con.Scan(&H, &W, &A, &B)
	M := make([][]bool, H)
	for i := 0; i < H; i++ {
		M[i] = make([]bool, W)
	}
	var rec func(h int, w int, a int, b int) int
	rec = func(h int, w int, a int, b int) int {
		for w < W && h < H && M[h][w] {
			w++
			if w >= W {
				w = 0
				h++
			}
		}
		if h >= H {
			return 1
		}
		M[h][w] = true
		var r int
		if b > 0 {
			r += rec(h, w, a, b-1)
		}
		if a > 0 && w < W-1 && !M[h][w+1] {
			M[h][w+1] = true
			r += rec(h, w, a-1, b)
			M[h][w+1] = false
		}
		if a > 0 && h < H-1 && !M[h+1][w] {
			M[h+1][w] = true
			r += rec(h, w, a-1, b)
			M[h+1][w] = false
		}
		M[h][w] = false
		return r
	}
	con.Println(rec(0, 0, A, B))
	return nil
}
