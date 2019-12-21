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

func (con *contest) scan(a ...interface{}) (int, error) {
	return fmt.Fscan(con.in, a...)
}
func (con *contest) scanln(a ...interface{}) (int, error) {
	return fmt.Fscanln(con.in, a...)
}
func (con *contest) scanf(f string, a ...interface{}) (int, error) {
	return fmt.Fscanf(con.in, f, a...)
}
func (con *contest) print(a ...interface{}) (int, error) {
	return fmt.Fprint(con.out, a...)
}
func (con *contest) println(a ...interface{}) (int, error) {
	return fmt.Fprintln(con.out, a...)
}
func (con *contest) printf(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(con.out, f, a...)
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	con := &contest{in: in, out: out}
	con.main()
}

type p struct{ r, c int }

func (con *contest) main() error {
	var H, W, N int
	con.scan(&H, &W, &N)
	M := map[p]int{}
	for i := 0; i < N; i++ {
		var a, b int
		con.scan(&a, &b)
		for r := a - 3; r < a; r++ {
			if r < 0 || r > H-3 {
				continue
			}
			for c := b - 3; c < b; c++ {
				if c < 0 || c > W-3 {
					continue
				}
				M[p{r, c}]++
			}
		}
	}
	A := make([]int, 10)
	for _, v := range M {
		A[v]++
	}
	A[0] = (H-2)*(W-2) - len(M)
	for _, a := range A {
		con.println(a)
	}
	return nil
}
