package main

// Solution for https://atcoder.jp/contests/abc182/tasks/abc182_c

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
	var s string
	con.Scan(&s)
	a := make([]int, 3)
	b := make([]int, 3)
	for i := 0; i < len(s); i++ {
		a[(s[i]-'0')%3]++
		b[(s[i]-'0')%3]++
	}
	for a[1] > 0 && a[2] > 0 {
		a[1]--
		a[2]--
	}
	for a[1] > 2 {
		a[1] -= 3
	}
	for a[2] > 2 {
		a[2] -= 3
	}
	x := a[1] + a[2]
	for b[1] > 2 {
		b[1] -= 3
	}
	for b[2] > 2 {
		b[2] -= 3
	}
	for b[1] > 0 && b[2] > 0 {
		b[1]--
		b[2]--
	}
	y := b[1] + b[2]
	if y < x {
		x = y
	}
	if x == len(s) {
		x = -1
	}
	con.Println(x)
	return nil
}
