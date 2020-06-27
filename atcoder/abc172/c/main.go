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
	var n, m, k int
	con.Scan(&n, &m, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		con.Scan(&b[i])
	}
	sum := 0
	max := 0
	p := 0
	q := 0
	for p < len(a) && sum+a[p] <= k {
		sum += a[p]
		p++
	}
	for q < len(b) && sum+b[q] <= k {
		sum += b[q]
		q++
	}
	max = p + q
	for p > 0 {
		p--
		sum -= a[p]
		for q < len(b) && sum+b[q] <= k {
			sum += b[q]
			q++
		}
		if p+q > max {
			max = p + q
		}
	}
	for q < len(b) && sum+b[q] <= k {
		sum += b[q]
		q++
	}
	if p+q > max {
		max = p + q
	}
	con.Println(max)
	return nil
}
