package main

// Solution for https://atcoder.jp/contests/abc186/tasks/abc186_c

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var n int
	con.Scan(&n)
	c := 0
	for i := 1; i <= n; i++ {
		if !strings.ContainsRune(fmt.Sprintf("%d", i), '7') && !strings.ContainsRune(fmt.Sprintf("%o", i), '7') {
			c++
		}
	}
	con.Println(c)
	return nil
}
