package main

import (
	"fmt"
	"sort"
)

var pq []int

func push(x int) {
	p := sort.Search(len(pq), func(i int) bool { return pq[i] <= x })
	if p == len(pq) {
		pq = append(pq, x)
		return
	}
	pq = append(pq[:p+1], pq[p:]...)
	pq[p] = x
}

func pop() int {
	if len(pq) == 0 {
		return 0
	}
	ans := pq[0]
	pq = pq[1:]
	return ans
}

func main() {
	pq = make([]int, 0, 5)
	push(1)
	fmt.Println(pq)
	push(3)
	fmt.Println(pq)
	push(2)
	fmt.Println(pq)
	push(1)
	fmt.Println(pq)
	fmt.Println(pop())
	fmt.Println(pq)
	fmt.Println(pop())
	fmt.Println(pq)
}
