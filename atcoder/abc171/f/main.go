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
	var k int
	var s string
	con.Scan(&k, &s)
	n := len(s)

	dp := make([][]int, 2)
	dp[0] = make([]int, n+1)
	dp[1] = make([]int, n+1)
	dp[0][0] = 1

	for i := 1; i <= n+k; i++ {
		p := i & 1
		q := p ^ 1
		for j := 0; j <= n; j++ {
			if j == n {
				dp[p][j] = dp[q][j] * 26
			} else {
				dp[p][j] = dp[q][j] * 25
			}
			dp[p][j] %= mod
			if j > 0 {
				dp[p][j] += dp[q][j-1]
				dp[p][j] %= mod
			}
		}
	}

	fmt.Println(dp[(n+k)&1][n])
	return nil
}

const mod = 1e9 + 7
