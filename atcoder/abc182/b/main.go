package main

// Solution for https://atcoder.jp/contests/abc182/tasks/abc182_b

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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
	}
	maxG := 0
	maxI := 0
	for i := 2; i <= 1000; i++ {
		g := 0
		for j := 0; j < n; j++ {
			if a[j]%i == 0 {
				g++
			}
		}
		if g >= maxG {
			maxG = g
			maxI = i
		}
	}
	con.Println(maxI)
	return nil
}
