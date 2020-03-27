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
	m := make([]int, 26)
	for i := 0; i < 26; i++ {
		m[i] = 100
	}
	for i := 0; i < n; i++ {
		var s string
		con.Scan(&s)
		d := make([]int, 26)
		for _, c := range s {
			d[c-'a']++
		}
		for j := 0; j < 26; j++ {
			if d[j] < m[j] {
				m[j] = d[j]
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < m[i]; j++ {
			con.Printf("%c", 'a'+i)
		}
	}
	con.Println()
	return nil
}
