package main

// Solution for https://atcoder.jp/contests/abc193/tasks/abc193_d

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

func valid(S string, K int) bool {
	H := make([]int, 10)
	for i := 0; i < len(S); i++ {
		H[S[i]-'0']++
	}
	for i := 1; i <= 9; i++ {
		if H[i] > K {
			return false
		}
	}
	return true
}

func score(S string) int {
	H := make([]int, 10)
	for i := 0; i < 5; i++ {
		H[S[i]-'0']++
	}
	s := 0
	for i := 1; i <= 9; i++ {
		m := 1
		for j := 0; j < H[i]; j++ {
			m *= 10
		}
		s += i * m
	}
	return s
}

func (con *contest) main() error {
	var K int
	var S, T string
	con.Scan(&K, &S, &T)
	H := make([]int, 10)
	for i := 1; i <= 9; i++ {
		H[i] = K
	}
	for _, a := range S[:4] + T[:4] {
		H[a-'0']--
	}
	M := 9*K - 8
	L := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			O := 0
			if i == j {
				if H[i] < 2 {
					continue
				}
				O = H[i] * (H[i] - 1)
			} else {
				if H[i] < 1 && H[j] < 1 {
					continue
				}
				O = H[i] * H[j]
			}
			s := S[:4] + fmt.Sprint(i)
			t := T[:4] + fmt.Sprint(j)
			if score(s) <= score(t) {
				continue
			}
			L += O
		}
	}
	con.Println(float64(L) / float64(M*(M-1)))
	return nil
}
