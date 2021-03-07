package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var testcases int
	con.Scan(&testcases)
	for testcase := 0; testcase < testcases; testcase++ {
		var N, K int
		var S string
		con.Scan(&N, &K, &S)
		if N%K != 0 || len(S) > N {
			con.Println(-1)
			continue
		}
		if len(S) < N {
			con.Println(strings.Repeat("a", N))
			continue
		}
		H := make([]int, 26)
		for _, d := range S {
			H[d-'a']++
		}
		var X, Y, I int
	loop:
		for I = N; I > 0; I-- {
			Y = int(S[I-1] - 'a')
			H[Y]--
			if I < N {
				Y++
			}
			for Y < 26 {
				H[Y]++
				X = modSum(H, K)
				//fmt.Printf("%d:%d:%v\n", I, X, H)
				if X <= N-I && (N-I-X)%K == 0 {
					break loop
				}
				H[Y]--
				Y++
			}
		}
		if I == 0 {
			con.Println(-1)
			continue
		}
		con.Printf("%s%c", S[:I-1], Y+'a')
		con.Print(strings.Repeat("a", N-I-X))
		for i := 0; i < 26; i++ {
			for H[i]%K != 0 {
				con.Printf("%c", 'a'+i)
				H[i]++
			}
		}
		con.Println()
	}
	return nil
}

func modSum(H []int, K int) int {
	s := 0
	for _, h := range H {
		m := h % K
		if m > 0 {
			s += K - m
		}
	}
	return s
}
