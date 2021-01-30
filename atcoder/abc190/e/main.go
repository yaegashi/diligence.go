package main

// Solution for https://atcoder.jp/contests/abc190/tasks/abc190_e

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
	var N, M, K int
	con.Scan(&N, &M)
	D := make([][]int, N)
	for i := 0; i < M; i++ {
		var a, b int
		con.Scan(&a, &b)
		a--
		b--
		D[a] = append(D[a], b)
		D[b] = append(D[b], a)
	}
	con.Scan(&K)
	C := make([]int, K)
	E := make([][]int, K)
	for i := 0; i < K; i++ {
		con.Scan(&C[i])
		C[i]--
		E[i] = make([]int, K)
	}
	if K == 1 {
		con.Println(1)
		return nil
	}
	for i := 0; i < K; i++ {
		V := make([]int, N)
		Q := make([]int, 0, N)
		Q = append(Q, C[i])
		for len(Q) > 0 {
			q := Q[0]
			for _, k := range D[q] {
				if k != C[i] && V[k] == 0 {
					V[k] = V[q] + 1
					Q = append(Q, k)
				}
			}
			Q = Q[1:]
		}
		for j := 0; j < K; j++ {
			if i == j {
				continue
			}
			q := C[j]
			if V[q] == 0 {
				con.Println(-1)
				return nil
			}
			E[i][j] = V[q]
		}
	}
	DP := make([][]int, 1<<K)
	for i := 0; i < 1<<K; i++ {
		DP[i] = make([]int, K)
	}
	var rec func(d int, v int) int
	rec = func(s int, v int) int {
		if DP[s][v] > 0 {
			return DP[s][v]
		}
		t := s ^ (1 << v)
		if t == 0 {
			return 0
		}
		minR := int(1e12)
		for u := 0; u < K; u++ {
			if t&(1<<u) != 0 {
				r := rec(t, u) + E[u][v]
				if r < minR {
					minR = r
				}
			}
		}
		DP[s][v] = minR
		return minR
	}
	minL := int(1e12)
	for i := 0; i < K; i++ {
		l := rec((1<<K)-1, i)
		if l < minL {
			minL = l
		}
	}
	con.Println(minL + 1)
	return nil
}
