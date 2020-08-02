package main

// Solution for https://atcoder.jp/contests/m-solutions2020/tasks/m_solutions2020_f

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

type pos struct{ x, y int }

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	U := []pos{}
	D := []pos{}
	R := []pos{}
	L := []pos{}
	for i := 0; i < N; i++ {
		var x, y int
		var u string
		con.Scan(&x, &y, &u)
		switch u {
		case "U":
			U = append(U, pos{x, y})
		case "D":
			D = append(D, pos{x, y})
		case "R":
			R = append(R, pos{x, y})
		case "L":
			L = append(L, pos{x, y})
		}
	}

	ht := 10000000
	i := 0
	j := 0

	sort.Slice(U, func(i, j int) bool {
		if U[i].x == U[j].x {
			return U[i].y < U[j].y
		} else {
			return U[i].x < U[j].x
		}
	})
	sort.Slice(D, func(i, j int) bool {
		if D[i].x == D[j].x {
			return D[i].y < D[j].y
		} else {
			return D[i].x < D[j].x
		}
	})
	i, j = 0, 0
	for i < len(U) && j < len(D) {
		if U[i].x < D[j].x {
			i++
		} else if U[i].x > D[j].x {
			j++
		} else if U[i].y < D[j].y {
			t := (D[j].y - U[i].y) * 5
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}

	sort.Slice(R, func(i, j int) bool {
		if R[i].y == R[j].y {
			return R[i].x < R[j].x
		} else {
			return R[i].y < R[j].y
		}
	})
	sort.Slice(L, func(i, j int) bool {
		if L[i].y == L[j].x {
			return L[i].x < L[j].x
		} else {
			return L[i].y < L[j].y
		}
	})
	i, j = 0, 0
	for i < len(R) && j < len(L) {
		if R[i].y < L[j].y {
			i++
		} else if R[i].y > L[j].y {
			j++
		} else if R[i].x < L[j].x {
			t := (L[j].x - R[i].x) * 5
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}

	sort.Slice(U, func(i, j int) bool {
		if U[i].x+U[i].y == U[j].x+U[j].y {
			return U[i].y < U[j].y
		} else {
			return U[i].x+U[i].y < U[j].x+U[j].y
		}
	})
	sort.Slice(D, func(i, j int) bool {
		if D[i].x+D[i].y == D[j].x+D[j].y {
			return D[i].y < D[j].y
		} else {
			return D[i].x+D[i].y < D[j].x+D[j].y
		}
	})
	sort.Slice(R, func(i, j int) bool {
		if R[i].x+R[i].y == R[j].x+R[j].y {
			return R[i].y < R[j].y
		} else {
			return R[i].x+R[i].y < R[j].x+R[j].y
		}
	})
	sort.Slice(L, func(i, j int) bool {
		if L[i].x+L[i].y == L[j].x+L[j].y {
			return L[i].y < L[j].y
		} else {
			return L[i].x+L[i].y < L[j].x+L[j].y
		}
	})
	i, j = 0, 0
	for i < len(U) && j < len(R) {
		if U[i].x+U[i].y < R[j].x+R[j].y {
			i++
		} else if U[i].x+U[i].y > R[j].x+R[j].y {
			j++
		} else if U[i].y < R[j].y {
			t := (R[j].y - U[i].y) * 10
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}
	i, j = 0, 0
	for i < len(L) && j < len(D) {
		if L[i].x+L[i].y < D[j].x+D[j].y {
			i++
		} else if L[i].x+L[i].y > D[j].x+D[j].y {
			j++
		} else if L[i].y < D[j].y {
			t := (D[j].y - L[i].y) * 10
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}

	sort.Slice(U, func(i, j int) bool {
		if U[i].x-U[i].y == U[j].x-U[j].y {
			return U[i].y < U[j].y
		} else {
			return U[i].x-U[i].y < U[j].x-U[j].y
		}
	})
	sort.Slice(D, func(i, j int) bool {
		if D[i].x-D[i].y == D[j].x-D[j].y {
			return D[i].y < D[j].y
		} else {
			return D[i].x-D[i].y < D[j].x-D[j].y
		}
	})
	sort.Slice(R, func(i, j int) bool {
		if R[i].x-R[i].y == R[j].x-R[j].y {
			return R[i].y < R[j].y
		} else {
			return R[i].x-R[i].y < R[j].x-R[j].y
		}
	})
	sort.Slice(L, func(i, j int) bool {
		if L[i].x-L[i].y == L[j].x-L[j].y {
			return L[i].y < L[j].y
		} else {
			return L[i].x-L[i].y < L[j].x-L[j].y
		}
	})
	i, j = 0, 0
	for i < len(R) && j < len(D) {
		if R[i].x-R[i].y < D[j].x-D[j].y {
			i++
		} else if R[i].x-R[i].y > D[j].x-D[j].y {
			j++
		} else if R[i].y < D[j].y {
			t := (D[j].y - R[i].y) * 10
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}
	i, j = 0, 0
	for i < len(U) && j < len(L) {
		if U[i].x-U[i].y < L[j].x-L[j].y {
			i++
		} else if U[i].x-U[i].y > L[j].x-L[j].y {
			j++
		} else if U[i].y < L[j].y {
			t := (L[j].y - U[i].y) * 10
			if t < ht {
				ht = t
			}
			i++
		} else {
			j++
		}
	}

	if ht == 10000000 {
		con.Println("SAFE")
	} else {
		con.Println(ht)
	}
	return nil
}
