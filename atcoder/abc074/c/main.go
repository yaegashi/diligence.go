package main

// Solution for https://atcoder.jp/contests/abc074/tasks/arc083_a

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
	var a, b, c, d, e, f int
	con.Scan(&a, &b, &c, &d, &e, &f)
	maxX, maxY := a, 0
	for i := 0; ; i++ {
		x := a * i
		if 100*x > f {
			break
		}
		for j := 0; ; j++ {
			x := x + b*j
			if 100*x > f {
				break
			}
			for k := 0; ; k++ {
				y := c * k
				if 100*x+y > f {
					break
				}
				for l := 0; ; l++ {
					y := y + d*l
					if 100*x+y > f {
						break
					}
					if y > x*e {
						break
					}
					if y*(maxX+maxY) > maxY*(x+y) {
						maxX = x
						maxY = y
					}
				}
			}
		}
	}
	con.Println(100*maxX+maxY, maxY)
	return nil
}
