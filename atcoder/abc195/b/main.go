package main

// Solution for https://atcoder.jp/contests/abc195/tasks/abc195_b

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
	var A, B, W int
	con.Scan(&A, &B, &W)
	W *= 1000
	minX := 1000000000
	maxX := 0
	for i := 1; ; i++ {
		if A*i > W {
			break
		}
		if B*i < W {
			continue
		}
		if i < minX {
			minX = i
		}
		if i > maxX {
			maxX = i
		}
	}
	if maxX == 0 {
		con.Println("UNSATISFIABLE")
	} else {
		con.Println(minX, maxX)
	}
	return nil
}
