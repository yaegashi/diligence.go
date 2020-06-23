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
		a := make([]int, n)
		for i := 0; i < n; i++ {
			con.Scan(&a[i])
		}
		sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		w := make([]int, k)
		for i := 0; i < k; i++ {
			con.Scan(&w[i])
		}
		sort.Ints(w)
		sum := 0
		for i := 0; i < k; i++ {
			sum += a[i]
		}
		base := 0
		lev := 1
		for base < k {
			pos := 0
			for base+pos < k && w[base+pos] == lev {
				sum += a[pos]
				pos++
			}
			a = a[k-base:]
			base += pos
			lev++
		}
		con.Println(sum)
	}
	return nil
}
