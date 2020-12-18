package main

// Solution for https://atcoder.jp/contests/abc175/tasks/abc175_d

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

func calc(A []int, K int) int {
	N := len(A)
	D := make([]int, N*2)
	S := 0
	for i := 0; i < N; i++ {
		D[i] = A[i]
		D[i+N] = A[i]
		S += A[i]
	}
	E := make([]int, N*2+1)
	for i := 1; i <= N*2; i++ {
		E[i] = E[i-1] + D[i-1]
	}
	X := make([]int, N+1)
	for i := 1; i <= N; i++ {
		X[i] = int(-1e10)
		for j := 0; j < N; j++ {
			if E[j+i]-E[j] > X[i] {
				X[i] = E[j+i] - E[j]
			}
		}
	}
	MAX := int(-1e10)
	for i := 1; i <= N && i <= K; i++ {
		if X[i] > MAX {
			MAX = X[i]
		}
	}
	if K > N && S > 0 {
		Y := K / N
		Y *= S
		M := K % N
		for i := 0; i <= M; i++ {
			if Y+X[i] > MAX {
				MAX = Y + X[i]
			}
		}
	}
	return MAX
}

func (con *contest) main() error {
	var N, K int
	con.Scan(&N, &K)
	P := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&P[i])
		P[i]--
	}
	for i := 0; i < N; i++ {
		con.Scan(&C[i])
	}
	V := make([]bool, N)
	MAX := int(-1e10)
	for i := 0; i < N; i++ {
		if V[i] {
			continue
		}
		x := i
		D := []int{}
		for !V[x] {
			V[x] = true
			D = append(D, C[x])
			x = P[x]
		}
		X := calc(D, K)
		if X > MAX {
			MAX = X
		}
	}
	con.Println(MAX)
	return nil
}
