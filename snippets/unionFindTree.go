package uftree

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
