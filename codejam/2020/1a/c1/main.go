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
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var R, C int
		con.Scan(&R, &C)
		M := make([][]int, R)
		E := make([][]bool, R)
		for i := 0; i < R; i++ {
			M[i] = make([]int, C)
			E[i] = make([]bool, C)
			for j := 0; j < C; j++ {
				con.Scan(&M[i][j])
			}
		}
		I := 0
		for {
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					if M[i][j] == 0 {
						E[i][j] = false
						continue
					}
					I += M[i][j]
					e := 0
					for k := i - 1; k >= 0; k-- {
						if M[k][j] > 0 {
							e += M[k][j]
							e -= M[i][j]
							break
						}
					}
					for k := i + 1; k < R; k++ {
						if M[k][j] > 0 {
							e += M[k][j]
							e -= M[i][j]
							break
						}
					}
					for k := j - 1; k >= 0; k-- {
						if M[i][k] > 0 {
							e += M[i][k]
							e -= M[i][j]
							break
						}
					}
					for k := j + 1; k < C; k++ {
						if M[i][k] > 0 {
							e += M[i][k]
							e -= M[i][j]
							break
						}
					}
					E[i][j] = e > 0
				}
			}
			done := true
			for i := 0; i < R; i++ {
				for j := 0; j < C; j++ {
					if E[i][j] {
						M[i][j] = 0
						done = false
					}
				}
			}
			if done {
				break
			}
		}
		con.Printf("Case #%d: %d\n", testcase, I)
	}
	return nil
}
