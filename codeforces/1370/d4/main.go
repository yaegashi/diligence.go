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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
	}
	hi := 1000000000
	lo := 1
	for hi > lo {
		mid := (hi + lo) / 2
		ok := false
		if !ok {
			j := 0
			for i := 0; i < n && j < k; i++ {
				if j%2 == 0 {
					if a[i] <= mid {
						j++
					}
				} else {
					j++
				}
			}
			ok = j == k
		}
		if !ok {
			j := 0
			for i := 0; i < n && j < k; i++ {
				if j%2 != 0 {
					if a[i] <= mid {
						j++
					}
				} else {
					j++
				}
			}
			ok = j == k
		}
		if ok {
			hi = mid
		} else if lo == mid {
			lo = mid + 1
		} else {
			lo = mid
		}
	}
	con.Println(lo)
	return nil
}
