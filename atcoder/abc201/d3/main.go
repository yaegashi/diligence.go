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
	}
	for h := H - 1; h >= 0; h-- {
		for w := W - 1; w >= 0; w-- {
			if h == 0 && w == 0 {
				continue
			}
			x := 1
			if M[h][w] == '-' {
				x = -1
			}
			if (w+h)%2 == 1 {
				D[h][w] += x
			} else {
				D[h][w] -= x
			}
			if h > 0 {
				D[h-1][w] = D[h][w]
			}
			if w > 0 {
				if (w+h)%2 == 1 {
					if D[h][w-1] > D[h][w] {
						D[h][w-1] = D[h][w]
					}
				} else {
					if D[h][w-1] < D[h][w] {
						D[h][w-1] = D[h][w]
					}
				}
			}
		}
	}
	Z := "Draw"
	if D[0][0] > 0 {
		Z = "Takahashi"
	}
	if D[0][0] < 0 {
		Z = "Aoki"
	}
	con.Println(Z)
	return nil
}
