package main

// Solution for https://atcoder.jp/contests/abc195/tasks/abc195_f

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

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}

func (con *contest) main() error {
	var A, B int
	con.Scan(&A, &B)
	C := B - A + 1
	Q := make([]int, C)
	R := 0
	for i := 0; i < C; i++ {
		for j := 0; j < len(primes); j++ {
			if (A+i)%primes[j] == 0 {
				Q[i] |= 1 << j
			}
		}
		if Q[i] == 0 {
			R++
		}
	}
	var dfs func(n int, factor int) int
	dfs = func(n int, factor int) int {
		if Q[n] == 0 || factor&Q[n] != 0 {
			return 0
		}
		factor |= Q[n]
		t := 1
		for i := n + 1; i < C; i++ {
			t += dfs(i, factor)
		}
		return t
	}
	T := 1
	for i := 0; i < C; i++ {
		T += dfs(i, 0)
	}
	T <<= R
	con.Println(T)
	return nil
}
