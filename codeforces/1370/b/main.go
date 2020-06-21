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

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 0; testcase < T; testcase++ {
		var n int
		con.Scan(&n)
		e := []int{}
		o := []int{}
		for i := 0; i < n+n; i++ {
			var a int
			con.Scan(&a)
			if a%2 == 0 {
				e = append(e, i+1)
			} else {
				o = append(o, i+1)
			}
		}
		if len(e)%2 == 0 && len(o)%2 == 0 {
			if len(e) > 0 {
				e = e[2:]
			} else {
				o = o[2:]
			}
		} else {
			e = e[1:]
			o = o[1:]
		}
		for i := 0; i < len(e)/2; i++ {
			con.Println(e[i*2], e[i*2+1])
		}
		for i := 0; i < len(o)/2; i++ {
			con.Println(o[i*2], o[i*2+1])
		}
	}
	return nil
}
