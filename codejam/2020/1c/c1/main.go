package main

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
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var N, D int
		con.Scan(&N, &D)
		A := make([]int, N)
		for i := 0; i < N; i++ {
			con.Scan(&A[i])
		}
		dp := make([][]int, N+1)
		for i := 0; i <= N; i++ {
			dp[i] = make([]int, D+1)
		}
		for i := 1; i <= D; i++ {
			dp[0][i] = D
		}
		minCut := D
		for d := 1; d <= D; d++ {
			done := map[int]bool{}
			for _, b := range A {
				if done[b] {
					continue
				}
				done[b] = true
				for i := 1; i <= N; i++ {
					for j := 1; j <= D; j++ {
						dp[i][j] = dp[i-1][j]
						s := A[i-1] * d / b
						if s == 0 {
							continue
						}
						for k := 1; k <= s && k <= j; k++ {
							c := k
							if k == s && A[i-1]*d%b == 0 {
								c--
							}
							if dp[i-1][j-k]+c < dp[i][j] {
								dp[i][j] = dp[i-1][j-k] + c
							}
						}
					}
				}
				if dp[N][D] < minCut {
					minCut = dp[N][D]
				}
			}
		}
		con.Printf("Case #%d: %d\n", testcase, minCut)
	}
	return nil
}
