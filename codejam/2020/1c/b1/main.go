package main

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var u int
		con.Scan(&u)
		h := make([][]string, 11)
		for i := 0; i < 10000; i++ {
			var m int
			var r string
			con.Scan(&m, &r)
			if m <= 10 {
				h[m] = append(h[m], r)
			}
		}
		var d string
		for _, ss := range h {
			for _, s := range ss {
				if len(s) > 1 {
					s = s[1:]
				}
				if !strings.Contains(d, s) {
					d += s
					continue
				}
			}
		}
		con.Printf("Case #%d: %s\n", testcase, d[9:]+d[:9])
	}
	return nil
}
