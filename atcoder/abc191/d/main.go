package main

// Solution for https://atcoder.jp/contests/abc191/tasks/abc191_d

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
	var X, Y, R float64
	con.Scan(&X, &Y, &R)
	XL := math.Floor(X - R)
	XR := math.Ceil(X + R)
	YT := math.Ceil(Y + R)
	YB := math.Floor(Y - R)
	RR := R * R
	a := 0
	for x := XL; x <= XR; x += 1.0 {
		var yt, yb float64
		xx := (x - X) * (x - X)
		ok := 0
		for y := YT; y >= Y; y -= 1.0 {
			if xx+(y-Y)*(y-Y) <= RR {
				ok++
				yt = y
				break
			}
		}
		for y := YB; y <= Y; y += 1.0 {
			if xx+(y-Y)*(y-Y) <= RR {
				ok++
				yb = y
				break
			}
		}
		if ok < 2 {
			continue
		}
		fmt.Println(x, a)
		a += int(yt-yb) + 1
	}
	con.Println(a)
	return nil
}
