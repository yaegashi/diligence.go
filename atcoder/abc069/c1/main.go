package main

// Solution for https://atcoder.jp/contests/abc069/tasks/arc080_a

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
	var n, a0, a2, a4 int
	con.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		con.Scan(&a)
		switch a % 4 {
		case 0:
			a4++
		case 2:
			a2++
		default:
			a0++
		}
	}
	possible := true
	if a0 > a4+1 {
		possible = false
	}
	if a2 > 0 {
		if a0 > a4 {
			possible = false
		}
	}
	if a2 == 1 {
		if a4 == 0 {
			possible = false
		}
	}
	if possible {
		con.Println("Yes")
	} else {
		con.Println("No")
	}
	return nil
}
