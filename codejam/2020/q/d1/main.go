package main

import (
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
	in := os.Stdin
	out := os.Stdout
	con := &contest{in: in, out: out}
	con.main()
}

func (con *contest) main() error {
	var T, B int
	con.Scan(&T, &B)
	for testcase := 1; testcase <= T; testcase++ {
		b := ""
		for i := 1; i <= B; i++ {
			con.Println(i)
			var s string
			con.Scan(&s)
			b += s
		}
		con.Println(b)
		var a string
		con.Scan(&a)
		if a == "N" {
			break
		}
	}
	return nil
}
