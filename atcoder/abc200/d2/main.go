package main

// Solution for https://atcoder.jp/contests/abc200/tasks/abc200_d

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

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	A := make([]int, N)
	B := make([][]int, 200)
	for i := 0; i < N; i++ {
		var a int
		con.Scan(&a)
		a %= 200
		A[i] = a
		B[a] = append(B[a], i+1)
	}
	/*
		for i := 0; i < 200; i++ {
			if len(B[i]) > 1 {
				con.Println("Yes")
				con.Println(1, B[i][0])
				con.Println(1, B[i][1])
				return nil
			}
		}
		if len(B[0]) > 0 {
			for i := 1; i < 200; i++ {
				if len(B[i]) > 0 {
					con.Println("Yes")
					con.Println(1, B[0][0])
					con.Println(2, B[0][0], B[i][0])
					return nil
				}
			}
			return nil
		}
	*/
	D := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		D[i] = make([][]int, 200)
	}
	for i := 1; i <= N; i++ {
		x := A[i-1]
		for j := 0; j < 200; j++ {
			k := (j - x + 200) % 200
			//fmt.Printf("%d:%d:%d:%d\n", i, j, x, k)
			if k == 0 && len(D[i-1][k]) > 0 {
				con.Println("Yes")
				con.dump([]int{i})
				con.dump(append(D[i-1][k], i))
				return nil
			}
			if (k == 0 || len(D[i-1][k]) > 0) && len(D[i-1][j]) > 0 {
				con.Println("Yes")
				con.dump(append(D[i-1][k], i))
				con.dump(D[i-1][j])
				return nil
			}
			if k == 0 || len(D[i-1][k]) > 0 {
				D[i][j] = make([]int, len(D[i-1][k]))
				copy(D[i][j], D[i-1][k])
				D[i][j] = append(D[i][j], i)
			} else {
				D[i][j] = make([]int, len(D[i-1][j]))
				copy(D[i][j], D[i-1][j])
			}
		}
	}
	con.Println("No")
	return nil
}

func (con *contest) dump(a []int) {
	sort.Ints(a)
	con.Print(len(a))
	for _, x := range a {
		con.Printf(" %d", x)
	}
	con.Println()
}
