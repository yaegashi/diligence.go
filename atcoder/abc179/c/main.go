package main

// Solution for https://atcoder.jp/contests/abc179/tasks/abc179_c

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
	s := make([]int, n+1)
	s[1] = 1
	for i := 2; i <= n; i++ {
		if s[i] == 0 {
			for j := i; j <= n; j += i {
				s[j] = i
			}
		}
	}
	var x int
	for c := 1; c < n; c++ {
		a := c
		m := map[int]int{}
		for a > 1 {
			m[s[a]]++
			a /= s[a]
		}
		y := 1
		for _, j := range m {
			//fmt.Printf("%d: %d ^ %d\n", c, i, j)
			y *= j + 1
		}
		x += y
	}
	con.Println(x)
	return nil
}
