package main

import (
	"container/heap"
	"fmt"
)

type priorityQueue []int

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i] > p[j] }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *priorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func main() {
	pq := make(priorityQueue, 0, 10)
	heap.Push(&pq, 1)
	heap.Push(&pq, 2)
	heap.Push(&pq, 1)
	heap.Push(&pq, 5)
	heap.Push(&pq, 8)
	heap.Push(&pq, 3)
	for len(pq) > 0 {
		x := heap.Pop(&pq).(int)
		fmt.Println(x)
	}
}
