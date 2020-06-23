package main

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
const N = 2000000

func (con *contest) main() error {
	s := []int{1, 0, 0}
	t := []int{0, 0, 0}
	a := make([]int, N+1)
	for i := 2; i <= N; i++ {
		s0 := s[0]
		s0 += s[1]
		s0 %= M
		s0 += s[1]
		s0 %= M
		s1 := s[0]
		s2 := s[1]
		s[0], s[1], s[2] = s0, s1, s2
		t[i%3] += s[2]
		t[i%3] %= M
		a[i] = t[i%3]
		a[i] %= M
		a[i] += t[i%3]
		a[i] %= M
		a[i] += t[i%3]
		a[i] %= M
		a[i] += t[i%3]
		a[i] %= M
	}
	var T int
	con.Scan(&T)
	for testcase := 0; testcase < T; testcase++ {
		var n int
		con.Scan(&n)
		con.Println(a[n])
	}
	return nil
}
