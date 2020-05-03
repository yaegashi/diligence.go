package main

type pqueue []int

func (pq pqueue) Len() int { return len(pq) }

func (pq pqueue) Less(i, j int) bool {
	return pq[i] < pq[j]
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
