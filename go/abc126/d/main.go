package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	edges  map[int]map[int]int
	colors []int
)

func dfs(v, d int) {
	if colors[v-1] > 0 {
		return
	}
	colors[v-1] = d + 2
	for u, w := range edges[v] {
		if w > 0 {
			dfs(u, (d+w)%2)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	edges = make(map[int]map[int]int, n)
	for i := 1; i <= n; i++ {
		edges[i] = make(map[int]int)
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n-1; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		w, _ := strconv.Atoi(sc.Text())
		edges[u][v] = w
		edges[v][u] = w
	}

	colors = make([]int, n)
	dfs(1, 0)

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range colors {
		wr.WriteString(strconv.Itoa(v-2) + "\n")
	}
	wr.Flush()
}
