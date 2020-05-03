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

func group(i int) int {
	switch i {
	case 1, 3, 5, 7, 8, 10, 12:
		return 1
	case 4, 6, 9, 11:
		return 2
	case 2:
		return 3
	}
	return 0
}

func (con *contest) main() error {
	var x, y int
	con.Scan(&x, &y)
	if group(x) == group(y) {
		con.Println("Yes")
	} else {
		con.Println("No")
	}
	return nil
}
