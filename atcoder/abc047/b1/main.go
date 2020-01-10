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
	var W, H, N int
	con.Scan(&W, &H, &N)
	M := make([][]bool, H)
	for i := range M {
		M[i] = make([]bool, W)
	}
	for i := 0; i < N; i++ {
		var x, y, a int
		con.Scan(&x, &y, &a)
		switch a {
		case 1:
			for j := 0; j < x; j++ {
				for k := 0; k < H; k++ {
					M[k][j] = true
				}
			}
		case 2:
			for j := x; j < W; j++ {
				for k := 0; k < H; k++ {
					M[k][j] = true
				}
			}
		case 3:
			for j := 0; j < y; j++ {
				for k := 0; k < W; k++ {
					M[j][k] = true
				}
			}
		case 4:
			for j := y; j < H; j++ {
				for k := 0; k < W; k++ {
					M[j][k] = true
				}
			}
		}
	}
	c := 0
	for i := range M {
		for _, a := range M[i] {
			if !a {
				c++
			}
		}
	}
	con.Println(c)
	return nil
}
