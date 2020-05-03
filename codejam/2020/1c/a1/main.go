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
		var x, y int
		var s string
		con.Scan(&x, &y, &s)
		a := 0
		for i, d := range s {
			switch d {
			case 'N':
				y++
			case 'S':
				y--
			case 'E':
				x++
			case 'W':
				x--
			}
			if abs(x)+abs(y) <= i+1 {
				a = i + 1
				break
			}
		}
		if a == 0 {
			con.Printf("Case #%d: IMPOSSIBLE\n", testcase)
		} else {
			con.Printf("Case #%d: %d\n", testcase, a)
		}
	}
	return nil
}
