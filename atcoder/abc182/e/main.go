package main

// Solution for https://atcoder.jp/contests/abc182/tasks/abc182_e

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

type obj struct {
	pos int
	lit bool
}

func (con *contest) main() error {
	var H, W, N, M int
	con.Scan(&H, &W, &N, &M)
	OH := make([][]obj, H)
	OV := make([][]obj, W)
	for i := 0; i < N; i++ {
		var a, b int
		con.Scan(&a, &b)
		OH[a-1] = append(OH[a-1], obj{b - 1, true})
		OV[b-1] = append(OV[b-1], obj{a - 1, true})
	}
	for i := 0; i < M; i++ {
		var a, b int
		con.Scan(&a, &b)
		OH[a-1] = append(OH[a-1], obj{b - 1, false})
		OV[b-1] = append(OV[b-1], obj{a - 1, false})
	}
	for h := 0; h < H; h++ {
		sort.Slice(OH[h], func(i, j int) bool { return OH[h][i].pos < OH[h][j].pos })
		//fmt.Println(OH[h])
	}
	for w := 0; w < W; w++ {
		sort.Slice(OV[w], func(i, j int) bool { return OV[w][i].pos < OV[w][j].pos })
		//fmt.Println(OV[w])
	}
	lit := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			j := sort.Search(len(OH[h]), func(i int) bool { return w <= OH[h][i].pos })
			k := sort.Search(len(OV[w]), func(i int) bool { return h <= OV[w][i].pos })
			//fmt.Println(h, w, j, k)
			if (j < len(OH[h]) && OH[h][j].pos == w && !OH[h][j].lit) || (k < len(OV[w]) && OV[w][k].pos == h && !OV[w][k].lit) {
				continue
			}
			if (0 < j && OH[h][j-1].lit) || (j < len(OH[h]) && OH[h][j].lit) || (0 < k && OV[w][k-1].lit) || (k < len(OV[w]) && OV[w][k].lit) {
				lit++
			}
		}
	}
	con.Println(lit)
	return nil
}
