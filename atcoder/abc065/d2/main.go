package main

import (
	"bufio"
	"container/heap"
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

type cost struct{ i, c int }

type pqueue []*cost

func (pq pqueue) Len() int { return len(pq) }

func (pq pqueue) Less(i, j int) bool {
	return pq[i].c < pq[j].c
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq pqueue) Peek() interface{} {
	return pq[0]
}

func (pq *pqueue) Push(x interface{}) {
	*pq = append(*pq, x.(*cost))
}

func (pq *pqueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func (con *contest) main() error {
	var n int
	con.Scan(&n)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		con.Scan(&x[i], &y[i])
	}
	res := 0
	pq := make(pqueue, n)
	pq[0] = &cost{0, 0}
	for i := 1; i < n; i++ {
		pq[i] = &cost{i, 1e18}
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		c := heap.Pop(&pq).(*cost)
		res += c.c
		u := c.i
		for i := 0; i < pq.Len(); i++ {
			v := pq[i].i
			cx := x[u] - x[v]
			if cx < 0 {
				cx = -cx
			}
			cy := y[u] - y[v]
			if cy < 0 {
				cy = -cy
			}
			if cx > cy {
				cx = cy
			}
			if cx < pq[i].c {
				pq[i].c = cx
			}
			heap.Fix(&pq, i)
		}
	}
	con.Println(res)
	return nil
}
