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
	h := make([]int, 10)
	for i := 0; i < len(s); i++ {
		h[s[i]-'0']++
	}
	for i := 1; i <= 9; i++ {
		if i%3 == 0 {
			h[i] = 0
			continue
		}
		for j := 1; j <= 9; j++ {
			if i == j {
				continue
			}
			if (i+j)%3 != 0 {
				continue
			}
			for h[i] > 0 && h[j] > 0 {
				h[i]--
				h[j]--
			}
		}
	}
	for i := 1; i <= 9; i++ {
		for h[i] > 3 {
			h[i] -= 3
		}
	}
	x := 0
	for i := 1; i <= 9; i++ {
		x += h[i]
	}
	if x == len(s) {
		x = -1
	}
	con.Println(x)
	return nil
}
