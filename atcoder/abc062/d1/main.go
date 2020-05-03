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

type PQLess []int

func (pq PQLess) Len() int { return len(pq) }

func (pq PQLess) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQLess) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQLess) Peek() interface{} {
	return pq[0]
}

func (pq *PQLess) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PQLess) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

type item struct {
	v, i int
}

type byOrder []item

func (b byOrder) Len() int           { return len(b) }
func (b byOrder) Less(i, j int) bool { return b[i].v < b[j].v }
func (b byOrder) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func (con *contest) main() error {
	var N int
	con.Scan(&N)

	p := make([]int, N)
	for i := 0; i < N; i++ {
		con.Scan(&p[i])
	}
	pPQ := (*PQLess)(&p)
	heap.Init(pPQ)

	q := make([]item, 2*N)
	r := make([]item, 2*N)
	for i := 0; i < 2*N; i++ {
		con.Scan(&r[i].v)
		r[i].i = i
	}
	sort.Sort(byOrder(r))
	for i := 0; i < 2*N; i++ {
		q[r[i].i].v = r[i].v
		q[r[i].i].i = i
	}

	pSum := 0
	qSum := 0
	for i := 0; i < N; i++ {
		pSum += p[i]
		qSum += r[i].v
	}

	s := N
	best := pSum - qSum
	for i := 0; i < N; i++ {
		if q[i].i < s {
			qSum -= q[i].v
			qSum += r[s].v
			r[s].v = 0
		}
		r[q[i].i].v = 0
		for s < len(r) && r[s].v == 0 {
			s++
		}
		if q[i].v > pPQ.Peek().(int) {
			pSum -= heap.Pop(pPQ).(int)
			pSum += q[i].v
			heap.Push(pPQ, q[i].v)
		}
		if pSum-qSum > best {
			best = pSum - qSum
		}
	}

	con.Println(best)

	return nil
}
