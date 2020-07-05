package main

// Solution for https://atcoder.jp/contests/abc173/tasks/abc173_e

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

const mod = 1e9 + 7

func (con *contest) main() error {
	var N, K int
	con.Scan(&N, &K)
	pos := make([]int, 0, N)
	neg := make([]int, 0, N)
	zero := 0
	for i := 0; i < N; i++ {
		var a int
		con.Scan(&a)
		if a > 0 {
			pos = append(pos, a)
		}
		if a < 0 {
			neg = append(neg, -a)
		}
		if a == 0 {
			zero++
		}
	}
	sort.Slice(pos, func(i, j int) bool { return pos[i] > pos[j] })
	sort.Slice(neg, func(i, j int) bool { return neg[i] > neg[j] })
	sum, i, j, k := 1, 0, 0, K
	for k > 1 {
		if i+1 < len(pos) && j+1 < len(neg) {
			if pos[i]*pos[i+1] > neg[j]*neg[j+1] {
				sum *= pos[i]
				sum %= mod
				sum *= pos[i+1]
				sum %= mod
				i += 2
			} else {
				sum *= neg[j]
				sum %= mod
				sum *= neg[j+1]
				sum %= mod
				j += 2
			}
		} else if i+1 < len(pos) {
			sum *= pos[i]
			sum %= mod
			sum *= pos[i+1]
			sum %= mod
			i += 2
		} else if j+1 < len(neg) {
			sum *= neg[j]
			sum %= mod
			sum *= neg[j+1]
			sum %= mod
			j += 2
		} else {
			break
		}
		k -= 2
	}
	if k == 0 {
		con.Println(sum)
		return nil
	}
	if k == 1 && i < len(pos) {
		sum *= pos[i]
		sum %= mod
		con.Println(sum)
		return nil
	}
	if zero > 0 {
		con.Println(0)
		return nil
	}
	for i, j := 0, len(pos)-1; i < j; i, j = i+1, j-1 {
		pos[i], pos[j] = pos[j], pos[i]
	}
	for i, j := 0, len(neg)-1; i < j; i, j = i+1, j-1 {
		neg[i], neg[j] = neg[j], neg[i]
	}
	sum, i, j, k = 1, 0, 0, K
	for k > 0 {
		if i < len(pos) && j < len(neg) {
			if pos[i] < neg[j] {
				sum *= pos[i]
				sum %= mod
				i++
			} else {
				sum *= neg[j]
				sum %= mod
				j++
			}
		} else if i < len(pos) {
			sum *= pos[i]
			sum %= mod
			i++
		} else if j < len(neg) {
			sum *= neg[j]
			sum %= mod
			j++
		}
		k--
	}
	con.Println(mod - sum)
	return nil
}
