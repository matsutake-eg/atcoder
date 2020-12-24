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
	ts := make(map[int][]int)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		ts[a] = append(ts[a], b)
	}

	pq := make(priorityQueue, 0, n)
	heap.Init(&pq)
	ans := 0
	for i := 1; i <= m; i++ {
		for _, v := range ts[i] {
			heap.Push(&pq, v)
		}
		if len(pq) > 0 {
			ans += heap.Pop(&pq).(int)
		}
	}
	fmt.Println(ans)
}
