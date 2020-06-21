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
	d0 := make([][]int, 2)
	d1 := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		d0[i] = make([]int, k+1)
		d1[i] = make([]bool, k+1)
	}
	for j := 1; j <= k; j++ {
		d0[0][j] = 1e9 + 1
		d1[0][j] = false
	}
	for i := 1; i <= n; i++ {
		p := i & 1
		q := p ^ 1
		for j := 1; j <= (k+1)/2; j++ {
			d0[p][j] = d0[q][j]
			d1[p][j] = false
			x := d0[q][j-1]
			if a[i] > x {
				x = a[i]
			}
			if !d1[q][j-1] && x < d0[p][j] {
				d0[p][j] = x
				d1[p][j] = true
			}
		}
	}
	m := d0[n&1][(k+1)/2]
	for j := 1; j <= k; j++ {
		d0[0][j] = 1e9 + 1
		d1[0][j] = true
	}
	for i := 1; i <= n; i++ {
		p := i & 1
		q := p ^ 1
		for j := 1; j <= k/2; j++ {
			d0[p][j] = d0[q][j]
			d1[p][j] = false
			x := d0[q][j-1]
			if a[i] > x {
				x = a[i]
			}
			if !d1[q][j-1] && x < d0[p][j] {
				d0[p][j] = x
				d1[p][j] = true
			}
		}
	}
	if d0[n&1][k/2] < m {
		m = d0[n&1][k/2]
	}
	con.Println(m)
	return nil
}
