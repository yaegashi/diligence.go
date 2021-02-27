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
	X := int(math.Round(Xf * 10000))
	Y := int(math.Round(Yf * 10000))
	R := int(math.Round(Rf * 10000))
	XL := floor(X - R)
	XR := ceil(X + R)
	RR := R * R
	a := 0
	for x := XL; x <= XR; x += 10000 {
		rr := RR - (x-X)*(x-X)
		if rr < 0 {
			continue
		}
		r := math.Sqrt(float64(rr))
		YT := int(math.Ceil(float64(Y) + r))
		YB := int(math.Floor(float64(Y) - r))
		var ok, yt, yb int
		for yt = ceil(YT + 100000); yt >= Y-100000; yt -= 10000 {
			if (yt-Y)*(yt-Y) <= rr {
				ok++
				break
			}
		}
		for yb = floor(YB - 100000); yb <= Y+100000; yb += 10000 {
			if (yb-Y)*(yb-Y) <= rr {
				ok++
				break
			}
		}
		if ok == 2 && yt >= yb {
			a += (yt-yb)/10000 + 1
		}
	}
	con.Println(a)
	return nil
}
