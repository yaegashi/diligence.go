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

const M = 1e9 + 7

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		var a int
		con.Scan(&a)
		A[a]++
	}
	ok := true
	if N%2 == 0 {
		for i := 0; i < N; i++ {
			if i%2 == 0 {
				if A[i] != 0 {
					ok = false
				}
			} else {
				if A[i] != 2 {
					ok = false
				}
			}
		}
	} else {
		if A[0] != 1 {
			ok = false
		}
		for i := 1; i < N; i++ {
			if i%2 == 0 {
				if A[i] != 2 {
					ok = false
				}
			} else {
				if A[i] != 0 {
					ok = false
				}
			}
		}
	}
	a := 0
	if ok {
		a = 1
		for i := 0; i < N/2; i++ {
			a = a + a
			a = a % M
		}
	}
	con.Println(a)
	return nil
}
