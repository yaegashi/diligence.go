package main

// Solution for https://atcoder.jp/contests/arc105/tasks/arc105_b

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

type pqueue []int

func (pq pqueue) Len() int { return len(pq) }

func (pq pqueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq pqueue) Peek() interface{} {
	return pq[0]
}

func (pq *pqueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *pqueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func (con *contest) main() error {
	var N int
	con.Scan(&N)
	pq := make(pqueue, N)
	min := int(1e9)
	for i := 0; i < N; i++ {
		con.Scan(&pq[i])
		if pq[i] < min {
			min = pq[i]
		}
	}
	heap.Init(&pq)
	for {
		X := heap.Pop(&pq).(int)
		for len(pq) > 0 && pq[0] == X {
			heap.Pop(&pq)
		}
		if X == min {
			con.Println(X)
			break
		}
		X -= min
		if X < min {
			min = X
		}
		heap.Push(&pq, X)
	}
	return nil
}
