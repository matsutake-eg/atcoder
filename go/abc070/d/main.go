package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

var n int
var G map[int]map[int]int

func initGraph() {
	G = make(map[int]map[int]int, n)
}

func add(from, to, weight int) {
	if _, ok := G[from]; !ok {
		G[from] = make(map[int]int)
	}
	G[from][to] = weight
}

var depth []int

func dfs(v, p, d int) {
	depth[v] = d
	for t, w := range G[v] {
		if t == p {
			continue
		}
		dfs(t, v, d+w)
	}
}

func main() {
	n = readInt()
	initGraph()
	for i := 0; i < n-1; i++ {
		a, b, c := readInt(), readInt(), readInt()
		add(a, b, c)
		add(b, a, c)
	}

	q, k := readInt(), readInt()
	depth = make([]int, n+1)
	dfs(k, -1, 0)
	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		x, y := readInt(), readInt()
		wr.WriteString(strconv.Itoa(depth[x]+depth[y]) + "\n")
	}
	wr.Flush()
}
