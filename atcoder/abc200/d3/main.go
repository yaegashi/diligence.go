package main

// Solution for https://atcoder.jp/contests/abc200/tasks/abc200_d

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
	var N int
	con.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&A[i])
		A[i] %= 200
	}
	if N > 8 {
		N = 8
	}
	B := make([]int, 200)
	for i := 1; i < 1<<N; i++ {
		S := 0
		for j := 0; j < N; j++ {
			if i&(1<<j) != 0 {
				S += A[j]
			}
		}
		S %= 200
		if B[S] > 0 {
			con.Println("Yes")
			con.dump(B[S])
			con.dump(i)
			return nil
		} else {
			B[S] = i
		}
	}
	con.Println("No")
	return nil
}

func (con *contest) dump(x int) {
	p := []interface{}{0}
	for i := 0; i < 8; i++ {
		if x&(1<<i) != 0 {
			p = append(p, i+1)
		}
	}
	p[0] = len(p) - 1
	con.Println(p...)
}
