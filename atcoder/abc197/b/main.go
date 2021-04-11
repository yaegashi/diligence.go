package main

// Solution for https://atcoder.jp/contests/abc197/tasks/abc197_b

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
	var H, W, X, Y int
	con.Scan(&H, &W, &X, &Y)
	X--
	Y--
	S := make([]string, H)
	for i := 0; i < H; i++ {
		con.Scan(&S[i])
	}
	C := 1
	for x, y := X-1, Y; x >= 0; x-- {
		if S[x][y] == '#' {
			break
		}
		C++
	}
	for x, y := X+1, Y; x < H; x++ {
		if S[x][y] == '#' {
			break
		}
		C++
	}
	for x, y := X, Y-1; y >= 0; y-- {
		if S[x][y] == '#' {
			break
		}
		C++
	}
	for x, y := X, Y+1; y < W; y++ {
		if S[x][y] == '#' {
			break
		}
		C++
	}
	con.Println(C)
	return nil
}
