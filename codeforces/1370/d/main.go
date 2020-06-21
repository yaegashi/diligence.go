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
	var n, k int
	con.Scan(&n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		con.Scan(&a[i])
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for j := 1; j <= k; j++ {
		dp[0][j] = 1e9 + 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			x := dp[i-1][j-1]
			if j%2 == 0 {
				if a[i] > x {
					x = a[i]
				}
			}
			if x < dp[i][j] {
				dp[i][j] = x
			}
		}
	}
	m := dp[n][k]
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			x := dp[i-1][j-1]
			if j%2 != 0 {
				if a[i] > x {
					x = a[i]
				}
			}
			if x < dp[i][j] {
				dp[i][j] = x
			}
		}
	}
	if dp[n][k] < m {
		m = dp[n][k]
	}
	con.Println(m)
	return nil
}
