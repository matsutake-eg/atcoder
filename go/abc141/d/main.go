package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
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
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	pq := make(priorityQueue, 0, n)
	ans := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		heap.Push(&pq, a)
		ans += a
	}

	for i := 0; i < m; i++ {
		a := heap.Pop(&pq).(int)
		ans -= a
		a /= 2
		ans += a
		heap.Push(&pq, a)
	}
	fmt.Println(ans)
}
