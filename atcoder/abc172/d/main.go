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
	var n int
	con.Scan(&n)
	sieve := make([]int, n+1)
	sieve[1] = 1
	for i := 2; i <= n; i++ {
		if sieve[i] != 0 {
			continue
		}
		for j := i; j <= n; j += i {
			sieve[j] = i
		}
	}
	sum := 0
	for i := 1; i <= n; i++ {
		x := i
		y := 1
		f := 0
		g := 0
		for x > 1 {
			if f != sieve[x] {
				y *= g + 1
				f = sieve[x]
				g = 0
			}
			x /= sieve[x]
			g++
		}
		y *= g + 1
		sum += y * i
	}
	con.Println(sum)
	return nil
}
