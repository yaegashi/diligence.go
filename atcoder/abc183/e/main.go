package main

// Solution for https://atcoder.jp/contests/abc183/tasks/abc183_e

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

const mod = 1e9 + 7

func (con *contest) main() error {
	var H, W int
	con.Scan(&H, &W)
	S := make([]string, H)
	for i := 0; i < H; i++ {
		con.Scan(&S[i])
	}
	DH := make([]int, H)
	DV := make([]int, W)
	DD := make([]int, H+W)
	T := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if S[h][w] == '#' {
				DV[w] = 0
				DH[h] = 0
				DD[w-h+H] = 0
				continue
			}
			if w == 0 && h == 0 {
				T = 1
			} else {
				T = DV[w] + DH[h]
				T %= mod
				T += DD[w-h+H]
				T %= mod
			}
			DV[w] += T
			DV[w] %= mod
			DH[h] += T
			DH[h] %= mod
			DD[w-h+H] += T
			DD[w-h+H] %= mod
		}
	}
	con.Println(T)
	return nil
}
