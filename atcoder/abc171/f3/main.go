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
	a := 0
	p25 := make([]int, k+1)
	p26 := make([]int, k+1)
	p25[0] = 1
	p26[0] = 1
	for i := 1; i <= k; i++ {
		p25[i] = p25[i-1] * 25
		p25[i] %= mod
		p26[i] = p26[i-1] * 26
		p26[i] %= mod
	}
	c := combmod(n+k-1, n-1)
	for i := 0; i <= k; i++ {
		x := c
		x *= p26[i]
		x %= mod
		x *= p25[k-i]
		x %= mod
		a += x
		a %= mod
		if i < k {
			c *= powmod(n+k-i-1, mod-2)
			c %= mod
			c *= k - i
			c %= mod
		}
	}
	fmt.Println(a)
	return nil
}

const mod = 1e9 + 7

func powmod(a, m int) int {
	b := 1
	for m > 0 {
		if m&1 != 0 {
			b = b * a % mod
		}
		a = a * a % mod
		m >>= 1
	}
	return b
}

func combmod(n, r int) int {
	c := 1
	for i := 0; i < r; i++ {
		c *= n - i
		c %= mod
		c *= powmod(i+1, mod-2)
		c %= mod
	}
	return c
}
