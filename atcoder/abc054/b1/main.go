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
	A := make([]string, N)
	B := make([]string, M)
	for i := 0; i < N; i++ {
		con.Scan(&A[i])
	}
	for i := 0; i < M; i++ {
		con.Scan(&B[i])
	}
	for i := 0; i+M <= N; i++ {
		for j := 0; j+M <= N; j++ {
			ok := true
			for k := 0; ok && k < M; k++ {
				for l := 0; ok && l < M; l++ {
					if A[i+k][j+l] != B[k][l] {
						ok = false
					}
				}
			}
			if ok {
				con.Println("Yes")
				return nil
			}
		}
	}
	con.Println("No")
	return nil
}
