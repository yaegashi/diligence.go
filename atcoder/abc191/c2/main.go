package main

// Solution for https://atcoder.jp/contests/abc191/tasks/abc191_c

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
	M := make([]string, H)
	for i := 0; i < H; i++ {
		con.Scan(&M[i])
	}
	a := 0
	for h := 1; h < H-1; h++ {
		for w := 1; w < W-1; w++ {
			if M[h][w] == '.' {
				continue
			}
			if M[h-1][w] == '.' && M[h][w-1] == '.' {
				a++
			}
			if M[h-1][w] == '.' && M[h][w+1] == '.' {
				a++
			}
			if M[h+1][w] == '.' && M[h][w-1] == '.' {
				a++
			}
			if M[h+1][w] == '.' && M[h][w+1] == '.' {
				a++
			}
			if M[h-1][w] == '#' && M[h][w-1] == '#' && M[h-1][w-1] == '.' {
				a++
			}
			if M[h-1][w] == '#' && M[h][w+1] == '#' && M[h-1][w+1] == '.' {
				a++
			}
			if M[h+1][w] == '#' && M[h][w-1] == '#' && M[h+1][w-1] == '.' {
				a++
			}
			if M[h+1][w] == '#' && M[h][w+1] == '#' && M[h+1][w+1] == '.' {
				a++
			}
		}
	}
	con.Println(a)
	return nil
}
