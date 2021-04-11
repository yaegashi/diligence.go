package main

// Solution for https://atcoder.jp/contests/abc198/tasks/abc198_d

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
	S := make([]string, 3)
	CM := make([]bool, 26)
	CC := 0
	for i := 0; i < 3; i++ {
		con.Scan(&S[i])
		for _, c := range S[i] {
			c -= 'a'
			if !CM[c] {
				CM[c] = true
				CC++
			}
		}
	}
	if CC > 10 {
		con.Println("UNSOLVABLE")
		return nil
	}
	C := make([]int, 26)
	for i := 0; i < 26; i++ {
		C[i] = -1
	}
	D := make([]int, 10)
	for i := 0; i < 10; i++ {
		D[i] = -1
	}
	ok := false
	calc := func() bool {
		M := make([]int, 3)
		for i := 0; i < 3; i++ {
			if C[S[i][0]-'a'] == 0 {
				return false
			}
			for _, c := range S[i] {
				M[i] = M[i]*10 + C[c-'a']
			}
		}
		if M[0]+M[1] == M[2] {
			con.Println(M[0])
			con.Println(M[1])
			con.Println(M[2])
			return true
		}
		return false
	}
	var rec func()
	rec = func() {
		var c int
		for c = 0; c < 26; c++ {
			if CM[c] && C[c] < 0 {
				break
			}
		}
		if c == 26 {
			ok = calc()
			return
		}
		for d := 0; !ok && d < 10; d++ {
			if D[d] < 0 {
				C[c] = d
				D[d] = c
				rec()
				C[c] = -1
				D[d] = -1
			}
		}
	}
	rec()
	if !ok {
		con.Println("UNSOLVABLE")
	}
	return nil
}
