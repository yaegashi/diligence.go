package main

// Solution for https://atcoder.jp/contests/abc186/tasks/abc186_b

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
	var H, W int
	con.Scan(&H, &W)
	A := make([]int, 101)
	minA := 100
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var a int
			con.Scan(&a)
			A[a]++
			if a < minA {
				minA = a
			}
		}
	}
	sumA := 0
	for a, x := range A {
		if a <= minA {
			continue
		}
		sumA += x * (a - minA)
	}
	con.Println(sumA)
	return nil
}
