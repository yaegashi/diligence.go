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

var mirrorTab = []int{0, 1, 5, -1, -1, 2, -1, -1, 8, -1}

func good(h int, m int, H int, M int) bool {
	D1 := mirrorTab[h/10]
	D2 := mirrorTab[h%10]
	D3 := mirrorTab[m/10]
	D4 := mirrorTab[m%10]
	if D1 < 0 || D2 < 0 || D3 < 0 || D4 < 0 {
		return false
	}
	h = D4*10 + D3
	m = D2*10 + D1
	return h < H && m < M
}

func (con *contest) main() error {
	var testcases int
	con.Scan(&testcases)
	for testcase := 0; testcase < testcases; testcase++ {
		var H, M int
		var S string
		con.Scan(&H, &M, &S)
		h := int(S[0]-'0')*10 + int(S[1]-'0')
		m := int(S[3]-'0')*10 + int(S[4]-'0')
		for !good(h, m, H, M) {
			m++
			if m >= M {
				m = 0
				h++
				if h >= H {
					h = 0
				}
			}
		}
		con.Printf("%02d:%02d\n", h, m)
	}
	return nil
}
