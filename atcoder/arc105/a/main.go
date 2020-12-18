package main

// Solution for https://atcoder.jp/contests/arc105/tasks/arc105_a

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
	a := make([]int, 4)
	con.Scan(&a[0], &a[1], &a[2], &a[3])
	ok := false
	for i := 0; i < 16; i++ {
		x := 0
		y := 0
		for j := 0; j < 4; j++ {
			if i&(1<<j) == 0 {
				x += a[j]
			} else {
				y += a[j]
			}
		}
		if x == y {
			ok = true
		}
	}
	if ok {
		con.Println("Yes")
	} else {
		con.Println("No")
	}
	return nil
}
