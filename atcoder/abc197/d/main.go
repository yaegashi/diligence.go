package main

// Solution for https://atcoder.jp/contests/abc197/tasks/abc197_d

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	var N, X0, Y0, X2, Y2 float64
	con.Scan(&N)
	con.Scan(&X0, &Y0, &X2, &Y2)
	cosT := math.Cos(2 * math.Pi / N)
	sinT := math.Sin(2 * math.Pi / N)
	xc := float64(X0+X2) / 2.0
	yc := float64(Y0+Y2) / 2.0
	x0 := X0 - xc
	y0 := Y0 - yc
	x1 := x0*cosT - y0*sinT + xc
	y1 := x0*sinT + y0*cosT + yc
	con.Println(x1, y1)
	return nil
}
