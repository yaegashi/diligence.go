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

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 0; testcase < T; testcase++ {
		var n, k int
		con.Scan(&n, &k)
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			con.Scan(&a[i])
		}
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		w := make([]int, k)
		for i := 0; i < k; i++ {
			con.Scan(&w[i])
		}
		sort.Slice(w, func(i, j int) bool { return w[i] > w[j] })
		var sum int64
		for i := n - k; i < n; i++ {
			sum += a[i]
		}
		j := 0
		for i, l := range w {
			if l == 1 {
				sum += a[n-k+i]
				continue
			}
			sum += a[j]
			j += l - 1
		}
		con.Println(sum)
	}
	return nil
}
