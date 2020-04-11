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
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var N int
		con.Scan(&N)
		I := 0
		B := 0
		for {
			M := N - I
			B = 0
			for i := 0; i < I; i++ {
				if M&1 == 1 {
					B += i + 1
				} else {
					B++
				}
				M >>= 1
			}
			if M == 0 {
				break
			}
			I++
		}
		con.Printf("Case #%d:\n", testcase)
		LR := false
		S := 0
		M := N - I
		for i := 1; i <= I; i++ {
			if M&1 == 1 {
				// scan
				if LR {
					for j := i; j >= 1; j-- {
						con.Println(i, j)
					}
				} else {
					for j := 1; j <= i; j++ {
						con.Println(i, j)
					}
				}
				LR = !LR
				S += 1 << uint(i-1)
			} else {
				// skip
				if LR {
					con.Println(i, i)
				} else {
					con.Println(i, 1)
				}
				S++
			}
			M >>= 1
		}
		for S < N {
			I++
			// skip
			if LR {
				con.Println(I, I)
			} else {
				con.Println(I, 1)
			}
			S++
		}
	}
	return nil
}
