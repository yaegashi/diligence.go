package main

// Solution for https://atcoder.jp/contests/hhkb2020/tasks/hhkb2020_d

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

const M = 1e9 + 7

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 0; testcase < T; testcase++ {
		var N, A, B int
		con.Scan(&N, &A, &B)
		S := 0
		NBA := N - B - A + 1
		if NBA > 0 {
			P := NBA * (NBA + 1)
			P %= M
			Q := (N - A + 1) * (N - B + 1)
			Q %= M
			R := P * Q
			R %= M
			S += R
			S %= M
			X := N - (A + B) + 1
			if X > 0 {
				Y := X * (X + 1)
				Y %= M
				Z := Q - Y
				if Z < 0 {
					Z += M
				}
				R := P * Z
				R %= M
				S += R
				S %= M
			}
		}
		con.Println(S)
	}
	return nil
}
