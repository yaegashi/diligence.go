package main

// Solution for https://atcoder.jp/contests/m-solutions2020/tasks/m_solutions2020_a

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
	var x int
	con.Scan(&x)
	q := 0
	if x < 600 {
		q = 8
	} else if x < 800 {
		q = 7
	} else if x < 1000 {
		q = 6
	} else if x < 1200 {
		q = 5
	} else if x < 1400 {
		q = 4
	} else if x < 1600 {
		q = 3
	} else if x < 1800 {
		q = 2
	} else {
		q = 1
	}
	con.Println(q)
	return nil
}
