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

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var N int
		con.Scan(&N)
		P := make([][]string, N)
		prefix := ""
		suffix := ""
		for i := 0; i < N; i++ {
			var p string
			con.Scan(&p)
			P[i] = strings.Split(p, "*")
			if len(P[i][0]) > len(prefix) {
				prefix = P[i][0]
			}
			if len(P[i][len(P[i])-1]) > len(suffix) {
				suffix = P[i][len(P[i])-1]
			}
		}
		ok := true
		for i := 0; ok && i < N; i++ {
			if len(prefix) > 0 && !strings.HasPrefix(prefix, P[i][0]) {
				ok = false
			}
			if len(suffix) > 0 && !strings.HasSuffix(suffix, P[i][len(P[i])-1]) {
				ok = false
			}
		}
		con.Printf("Case #%d: ", testcase)
		if ok {
			con.Println(prefix + suffix)
		} else {
			con.Println("*")
		}
	}
	return nil
}
