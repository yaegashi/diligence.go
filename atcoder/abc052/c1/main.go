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
	var N int
	con.Scan(&N)
	sieve := make([]int, N+1)
	for i := 2; i <= N; i++ {
		if sieve[i] != 0 {
			continue
		}
		for j := i; j <= N; j += i {
			sieve[j] = i
		}
	}
	primes := make([]int, N+1)
	for i := 2; i <= N; i++ {
		a := i
		for a > 1 {
			primes[sieve[a]]++
			a /= sieve[a]
		}
	}
	answer := 1
	for i := 2; i <= N; i++ {
		answer *= primes[i] + 1
		answer %= 1e9 + 7
	}
	con.Println(answer)
	return nil
}
