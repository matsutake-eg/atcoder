package main

import (
	"container/heap"
	"fmt"
)

type data struct{ a, b, c int }
type priorityQueue []*data

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].a > p[j].a }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(*data)) }
func (p *priorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func main() {
	pq := make(priorityQueue, 0, 10)
	heap.Push(&pq, &data{1, 2, 3})
	heap.Push(&pq, &data{2, 3, 4})
	heap.Push(&pq, &data{1, 3, 4})
	heap.Push(&pq, &data{3, 3, 4})
	heap.Push(&pq, &data{5, 3, 4})
	for len(pq) > 0 {
		x := heap.Pop(&pq).(*data)
		fmt.Println(x.a, x.b, x.c)
	}
}
