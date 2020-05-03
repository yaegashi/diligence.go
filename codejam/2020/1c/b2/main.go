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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type occurrence struct {
	x rune
	y int
}

type byOccurrence []occurrence

func (b byOccurrence) Len() int           { return len(b) }
func (b byOccurrence) Less(i, j int) bool { return b[i].y > b[j].y }
func (b byOccurrence) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func (con *contest) main() error {
	var T int
	con.Scan(&T)
	for testcase := 1; testcase <= T; testcase++ {
		var u int
		con.Scan(&u)
		c := map[rune]bool{}
		h := map[rune]int{}
		for i := 0; i < 10000; i++ {
			var m int
			var r string
			con.Scan(&m, &r)
			h[rune(r[0])]++
			if len(c) < 10 {
				for _, j := range r {
					c[j] = true
				}
			}
		}
		o := []occurrence{}
		for i, j := range h {
			o = append(o, occurrence{i, j})
		}
		sort.Sort(byOccurrence(o))
		d := ""
		for _, i := range o {
			d += string(i.x)
		}
		for i := range c {
			if h[i] == 0 {
				d = string(i) + d
			}
		}
		con.Printf("Case #%d: %s\n", testcase, d)
	}
	return nil
}
