package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

type P struct{ first, second int }

type ByP []P

func (b ByP) Len() int           { return len(b) }
func (b ByP) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByP) Less(i, j int) bool { return b[i].first < b[j].first }

func (con *contest) main() error {
	var N, K int
	con.Scan(&N, &K)
	A := []P{}
	for i := 0; i < N; i++ {
		var a, b int
		con.Scan(&a, &b)
		A = append(A, P{a, b})
	}
	sort.Sort(ByP(A))
	for _, a := range A {
		if K <= a.second {
			con.Println(a.first)
			break
		}
		K -= a.second
	}
	return nil
}
