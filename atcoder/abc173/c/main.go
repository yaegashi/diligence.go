package main

// Solution for https://atcoder.jp/contests/abc173/tasks/abc173_c

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
	var H, W, K int
	con.Scan(&H, &W, &K)
	M := make([]string, H)
	for h := 0; h < H; h++ {
		con.Scan(&M[h])
	}
	c := 0
	for i := 0; i < 1<<H; i++ {
		for j := 0; j < 1<<W; j++ {
			k := 0
			for h := 0; h < H; h++ {
				if i&(1<<h) != 0 {
					continue
				}
				for w := 0; w < W; w++ {
					if j&(1<<w) != 0 {
						continue
					}
					if M[h][w] == '#' {
						k++
					}
				}
			}
			if k == K {
				c++
			}
		}
	}
	con.Println(c)
	return nil
}
