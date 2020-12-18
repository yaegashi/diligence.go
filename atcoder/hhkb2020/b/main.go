package main

// Solution for https://atcoder.jp/contests/hhkb2020/tasks/hhkb2020_b

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
	var H, W int
	con.Scan(&H, &W)
	R := make([]string, H)
	for h := 0; h < H; h++ {
		con.Scan(&R[h])
	}
	c := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if R[h][w] == '.' && w < W-1 && R[h][w+1] == '.' {
				c++

			}
			if R[h][w] == '.' && h < H-1 && R[h+1][w] == '.' {
				c++
			}
		}
	}
	con.Println(c)
	return nil
}
