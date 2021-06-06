package main

// Solution for https://atcoder.jp/contests/abc201/tasks/abc201_d

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
	for h := 0; h < H; h++ {
		con.Scan(&M[h])
	}
	D := make([][]int, H)
	for h := 0; h < H; h++ {
		D[h] = make([]int, W)
		for w := 0; w < W; w++ {
			if h == 0 {
				continue
			}
			X := 0
			if (w+h)%2 == 1 {
				if h > 0 {
					X = D[h-1][w]
				}
				if w > 0 && D[h][w-1] > X {
					X = D[h][w-1]
				}
				if M[h][w] == '+' {
					X++
				} else {
					X--
				}
			} else {
				if h > 0 {
					X = D[h-1][w]
				}
				if w > 0 && D[h][w-1] < X {
					X = D[h][w-1]
				}
				if M[h][w] == '+' {
					X--
				} else {
					X++
				}
			}
			D[h][w] = X
		}
	}
	Z := "Draw"
	if D[H-1][W-1] > 0 {
		Z = "Takahashi"
	}
	if D[H-1][W-1] < 0 {
		Z = "Aoki"
	}
	con.Println(Z)
	return nil
}
