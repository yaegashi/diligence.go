package main

import (
	"bufio"
	"container/heap"
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

type cost struct{ i, c int }

type pqueue []edge

func (pq pqueue) Len() int { return len(pq) }

func (pq pqueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq pqueue) Peek() interface{} {
	return pq[0]
}

func (pq *pqueue) Push(x interface{}) {
	*pq = append(*pq, x.(edge))
}

func (pq *pqueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

type byX struct{ x, y, z []int }

func (b byX) Len() int           { return len(b.z) }
func (b byX) Less(i, j int) bool { return b.x[b.z[i]] < b.x[b.z[j]] }
func (b byX) Swap(i, j int)      { b.z[i], b.z[j] = b.z[j], b.z[i] }

type byY struct{ x, y, z []int }

func (b byY) Len() int           { return len(b.z) }
func (b byY) Less(i, j int) bool { return b.y[b.z[i]] < b.y[b.z[j]] }
func (b byY) Swap(i, j int)      { b.z[i], b.z[j] = b.z[j], b.z[i] }

type edge struct{ to, cost int }

func (con *contest) main() error {
	var n int
	con.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)
	v := make([][]edge, n)
	for i := 0; i < n; i++ {
		con.Scan(&x[i], &y[i])
		z[i] = i
	}
	sort.Sort(byX{x, y, z})
	for i := 0; i < n-1; i++ {
		c := x[z[i+1]] - x[z[i]]
		v[z[i]] = append(v[z[i]], edge{z[i+1], c})
		v[z[i+1]] = append(v[z[i+1]], edge{z[i], c})
	}
	sort.Sort(byY{x, y, z})
	for i := 0; i < n-1; i++ {
		c := y[z[i+1]] - y[z[i]]
		v[z[i]] = append(v[z[i]], edge{z[i+1], c})
		v[z[i+1]] = append(v[z[i+1]], edge{z[i], c})
	}
	res := 0
	used := make([]bool, n)
	pq := make(pqueue, 0, n*4)
	heap.Push(&pq, edge{0, 0})
	for pq.Len() > 0 {
		e := heap.Pop(&pq).(edge)
		if used[e.to] {
			continue
		}
		used[e.to] = true
		res += e.cost
		for _, x := range v[e.to] {
			if used[x.to] {
				continue
			}
			heap.Push(&pq, x)
		}
	}
	con.Println(res)
	return nil
}
