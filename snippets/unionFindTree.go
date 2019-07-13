package main

import "fmt"

var (
	par  []int
	rank []int
)

func initTree(x int) {
	par = make([]int, x+1)
	rank = make([]int, x+1)
	for i := range par {
		par[i] = i
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	}

	par[x] = find(par[x])
	return par[x]
}

func union(x, y int) {
	px, py := find(x), find(y)
	if px == py {
		return
	}

	if rank[px] < rank[py] {
		par[px] = py
		return
	} else if rank[px] == rank[py] {
		rank[px]++
	}
	par[py] = px
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	initTree(n)
	var x, y int
	for i := 0; i < k; i++ {
		fmt.Scan(&x, &y)
		union(x, y)
	}

	for i := 1; i <= n; i++ {
		fmt.Println("node:", i, "parent:", find(i))
	}
}
