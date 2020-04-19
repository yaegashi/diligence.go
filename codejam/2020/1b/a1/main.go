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
		con.Scan(&x, &y)
		absXY := abs(x) + abs(y)
		if absXY&1 == 0 {
			con.Printf("Case #%d: IMPOSSIBLE\n", testcase)
			continue
		}
		sign := []bool{}
		for a := absXY; a > 0; a >>= 1 {
			if a&1 != 0 {
				sign = append(sign, true)
			} else {
				sign = append(sign, false)
			}
		}
		for i := range sign {
			if (absXY>>uint(i))&1 == 0 {
				sign[i] = true
				sign[i-1] = false
			}
		}
		d := ""
		for i := 0; i < 1<<uint(len(sign)); i++ {
			sumX := 0
			sumY := 0
			for j, s := range sign {
				a := 1 << uint(j)
				if (i>>uint(j))&1 == 0 {
					if s {
						sumX += a
					} else {
						sumX -= a
					}
				} else {
					if s {
						sumY += a
					} else {
						sumY -= a
					}
				}
			}
			if abs(sumX) != abs(x) || abs(sumY) != abs(y) {
				continue
			}
			for j, s := range sign {
				if (i>>uint(j))&1 == 0 {
					if sumX != x {
						s = !s
					}
					if s {
						d += "E"
					} else {
						d += "W"
					}
				} else {
					if sumY != y {
						s = !s
					}
					if s {
						d += "N"
					} else {
						d += "S"
					}
				}
			}
			break
		}
		if len(d) == 0 {
			d = "IMPOSSIBLE"
		}
		con.Printf("Case #%d: %s\n", testcase, d)
	}
	return nil
}
