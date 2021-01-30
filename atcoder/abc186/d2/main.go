package main

// Solution for https://atcoder.jp/contests/abc186/tasks/abc186_d

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	A := map[int]int{}
	S := 0
	for i := 0; i < N; i++ {
		var a int
		con.Scan(&a)
		A[a]++
		S += a
	}
	var X []int
	for i := range A {
		X = append(X, i)
	}
	sort.Ints(X)
	Z := 0
	for _, x := range X {
		S -= x * A[x]
		N -= A[x]
		Z += (S - N*x) * A[x]
	}
	con.Println(Z)
	return nil
}
