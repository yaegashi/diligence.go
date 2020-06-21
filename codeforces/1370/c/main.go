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
		var a bool
		if n == 1 {
			a = false
		} else if n == 2 {
			a = true
		} else if n%2 != 0 {
			a = true
		} else {
			t := 0
			for n%2 == 0 {
				t++
				n /= 2
			}
			if n == 1 {
				a = t == 1
			} else if t != 1 {
				a = true
			} else {
				for i := 3; i*i <= n; i++ {
					if n%i == 0 {
						a = true
						break
					}
				}
			}
		}
		if a {
			con.Println("Ashishgup")
		} else {
			con.Println("FastestFinger")
		}
	}
	return nil
}
