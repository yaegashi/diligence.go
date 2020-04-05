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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&a[i])
	}
	var sumP, sumN, movP, movN int
	for i := 0; i < n; i++ {
		sumP += a[i]
		sumN += a[i]
		if i&1 == 0 {
			if sumP <= 0 {
				movP += 1 - sumP
				sumP = 1
			}
			if sumN >= 0 {
				movN += sumN + 1
				sumN = -1
			}
		} else {
			if sumP >= 0 {
				movP += sumP + 1
				sumP = -1
			}
			if sumN <= 0 {
				movN += 1 - sumN
				sumN = 1
			}
		}
	}
	if movP < movN {
		con.Println(movP)
	} else {
		con.Println(movN)
	}
	return nil
}
