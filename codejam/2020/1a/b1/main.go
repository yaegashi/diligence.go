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
	for testcase := 1; testcase <= T; testcase++ {
		var N int
		con.Scan(&N)
		con.Printf("Case #%d:\n", testcase)
		if N <= 500 {
			for i := 1; i <= N; i++ {
				con.Printf("%d 1\n", i)
			}
		} else {
			N--
			con.Println("1 1")
			i := 1
			for i <= N {
				con.Printf("%d 2\n", i+1)
				N -= i
				i++
			}
			for N > 0 {
				con.Printf("%d 1\n", i)
				N--
				i--
			}
		}
	}
	return nil
}
