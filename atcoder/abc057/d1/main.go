package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func (con *contest) comb(m, n int) int {
	p := [][]int{make([]int, m+1), make([]int, m+1)}
	p[0][0] = 1
	p[1][0] = 1
	for i := 1; i <= m; i++ {
		a := i & 1
		b := (i - 1) & 1
		for j := 1; j <= i; j++ {
			p[a][j] = p[b][j] + p[b][j-1]
		}
	}
	return p[m&1][n]
}

func (con *contest) main() error {
	var N, A, B int
	con.Scan(&N, &A, &B)
	v := make([]int, N)
	s := make([]int, N+1)
	for i := 0; i < N; i++ {
		con.Scan(&v[i])
	}
	sort.Ints(v)
	for i := 1; i <= N; i++ {
		s[i] = s[i-1] + v[N-i]
	}
	maxa := 0.0
	maxi := 0
	for i := A; i <= B; i++ {
		a := float64(s[i]) / float64(i)
		if a > maxa {
			maxa = a
			maxi = i
		}
	}
	con.Println(maxa)
	c := 0
	for i := A; i <= B; i++ {
		if s[i]*maxi == s[maxi]*i {
			var m, n int
			for j := 0; j < N; j++ {
				if v[j] == v[N-i] {
					m++
					if j >= N-i {
						n++
					}
				}
			}
			c += con.comb(m, n)
		}
	}
	con.Println(c)
	return nil
}
