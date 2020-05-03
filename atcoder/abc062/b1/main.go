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
	var h, w int
	con.Scan(&h, &w)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		con.Scan(&s[i])
	}
	for i := 0; i < w+2; i++ {
		con.Print("#")
	}
	con.Println()
	for i := 0; i < h; i++ {
		con.Printf("#%s#\n", s[i])
	}
	for i := 0; i < w+2; i++ {
		con.Print("#")
	}
	con.Println()
	return nil
}
