package main

// Solution for https://atcoder.jp/contests/abc191/tasks/abc191_d

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

func floor(x int) int {
	if x >= 0 {
		return x - x%10000
	}
	m := (-x) % 10000
	if m > 0 {
		x -= 10000 - m
	}
	return x
}

func ceil(x int) int {
	if x >= 0 {
		m := x % 10000
		if m > 0 {
			x += 10000 - m
		}
		return x
	}
	return x + (-x)%10000
}

func (con *contest) main() error {
	var Xf, Yf, Rf float64
	con.Scan(&Xf, &Yf, &Rf)
	X := int(Xf * 10000)
	Y := int(Yf * 10000)
	R := int(Rf * 10000)
	XL := floor(X - R)
	XR := ceil(X + R)
	RR := R * R
	a := 0
	for x := XL; x <= XR; x += 10000 {
		r := RR - (x-X)*(x-X)
		if r < 0 {
			continue
		}
		var yyt, yyb int
		{
			yt := Y + R
			yb := Y
			ym := 0
			for yb < yt-1 {
				ym = (yt + yb) / 2
				if (ym-Y)*(ym-Y) == r {
					break
				}
				if (ym-Y)*(ym-Y) <= r {
					yb = ym
				} else {
					yt = ym
				}
			}
			yyt = floor(ym)
		}
		{
			yt := Y
			yb := Y - R
			ym := 0
			for yb < yt-1 {
				ym = (yt + yb) / 2
				if (ym-Y)*(ym-Y) == r {
					break
				}
				if (ym-Y)*(ym-Y) <= r {
					yt = ym
				} else {
					yb = ym
				}
			}
			yyb = ceil(ym)
		}
		if yyt >= yyb {
			a += (yyt-yyb)/10000 + 1
		}
	}
	con.Println(a)
	return nil
}
