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
	var N, W int
	con.Scan(&N, &W)
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&w[i], &v[i])
	}
	w0 := w[0]
	for i := 0; i < N; i++ {
		w[i] -= w0
	}
	dp := make([][][]int, 2)
	dp[0] = make([][]int, N+1)
	dp[1] = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[0][i] = make([]int, N*3+1)
		dp[1][i] = make([]int, N*3+1)
		for j := 0; j <= N*3; j++ {
			dp[0][i][j] = -1
			dp[1][i][j] = -1
		}
	}
	dp[0][0][0] = 0
	dp[1][0][0] = 0
	maxV := 0
	for i := 1; i <= N; i++ {
		w1 := w[i-1]
		v1 := v[i-1]
		d1 := dp[i&1]
		d0 := dp[(i-1)&1]
		for j := 1; j <= N; j++ {
			for k := 0; k <= N*3; k++ {
				if j*w0+k <= W && k >= w1 && d0[j-1][k-w1] >= 0 && d0[j-1][k-w1]+v1 > d0[j][k] {
					d1[j][k] = d0[j-1][k-w1] + v1
					if d1[j][k] > maxV {
						maxV = d1[j][k]
					}
				} else {
					d1[j][k] = d0[j][k]
				}
			}
		}
	}
	con.Println(maxV)
	return nil
}
