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
	var N, M int
	con.Scan(&N, &M)
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, M)
	d := make([]int, M)
	for i := 0; i < N; i++ {
		con.Scan(&a[i], &b[i])
	}
	for i := 0; i < M; i++ {
		con.Scan(&c[i], &d[i])
	}
	for i := 0; i < N; i++ {
		mind := int(1e9)
		minj := -1
		for j := 0; j < M; j++ {
			x := a[i] - c[j]
			if x < 0 {
				x = -x
			}
			y := b[i] - d[j]
			if y < 0 {
				y = -y
			}
			if x+y < mind {
				mind = x + y
				minj = j
			}
		}
		con.Println(minj + 1)
	}
	return nil
}
