package main

// Solution for https://atcoder.jp/contests/abc201/tasks/abc201_c

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
	var S string
	con.Scan(&S)
	C := 0
	for X := 0; X < 10000; X++ {
		D := make([]int, 10)
		D[X%10]++
		D[(X/10)%10]++
		D[(X/100)%10]++
		D[(X/1000)%10]++
		ok := true
		for i := 0; i < 10; i++ {
			switch S[i] {
			case 'o':
				if D[i] == 0 {
					ok = false
				}
			case 'x':
				if D[i] != 0 {
					ok = false
				}
			}
		}
		if ok {
			C++
		}
	}
	con.Println(C)
	return nil
}
