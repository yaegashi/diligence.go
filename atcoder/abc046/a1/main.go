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

func (con *contest) scan(a ...interface{}) (int, error) {
	return fmt.Fscan(con.in, a...)
}
func (con *contest) scanln(a ...interface{}) (int, error) {
	return fmt.Fscanln(con.in, a...)
}
func (con *contest) scanf(f string, a ...interface{}) (int, error) {
	return fmt.Fscanf(con.in, f, a...)
}
func (con *contest) print(a ...interface{}) (int, error) {
	return fmt.Fprint(con.out, a...)
}
func (con *contest) println(a ...interface{}) (int, error) {
	return fmt.Fprintln(con.out, a...)
}
func (con *contest) printf(f string, a ...interface{}) (int, error) {
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
	a := 0
	m := map[int]bool{}
	con.scan(&a)
	m[a] = true
	con.scan(&a)
	m[a] = true
	con.scan(&a)
	m[a] = true
	con.println(len(m))
	return nil
}
