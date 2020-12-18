package main

// Solution for https://atcoder.jp/contests/abc183/tasks/abc183_d

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

const max = 200000

func (con *contest) main() error {
	var N, W int
	con.Scan(&N, &W)
	T := make([]int, max+1)
	for i := 0; i < N; i++ {
		var s, t, p int
		con.Scan(&s, &t, &p)
		T[s] += p
		T[t] -= p
	}
	ok := true
	w := 0
	for i := 0; i <= max; i++ {
		w += T[i]
		if w > W {
			ok = false
		}
	}
	if ok {
		con.Println("Yes")
	} else {
		con.Println("No")
	}
	return nil
}
