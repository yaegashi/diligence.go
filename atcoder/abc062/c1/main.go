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

func split1(h, w int) int {
	if w < 3 {
		return h * w
	}
	a := w / 3
	b := (w - a) / 2
	c := w - a - b
	return (c - a) * h
}

func split2(h, w int) int {
	a := h / 2
	b := h - a
	m := h * w
	for i := 1; i < w; i++ {
		c := []int{i * h, (w - i) * a, (w - i) * b}
		sort.Ints(c)
		if c[2]-c[0] < m {
			m = c[2] - c[0]
		}
	}
	return m
}

func (con *contest) main() error {
	var h, w int
	con.Scan(&h, &w)
	a := split1(h, w)
	b := split1(w, h)
	if b < a {
		a = b
	}
	b = split2(h, w)
	if b < a {
		a = b
	}
	b = split2(w, h)
	if b < a {
		a = b
	}
	con.Println(a)
	return nil
}
