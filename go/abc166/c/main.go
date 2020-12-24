package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

var graph map[int]map[int]bool

func initGraph(x int) {
	graph = make(map[int]map[int]bool, x)
}

func add(from, to int) {
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[int]bool)
	}
	graph[from][to] = true
}

func main() {
	n, m := scanInt(), scanInt()
	h := make([]int, n)
	for i := range h {
		h[i] = scanInt()
	}

	initGraph(n)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		add(a, b)
		add(b, a)
	}

	ans := 0
	for i := 0; i < n; i++ {
		isTop := true
		for k := range graph[i+1] {
			if h[k-1] >= h[i] {
				isTop = false
				break
			}
		}
		if isTop {
			ans++
		}
	}
	fmt.Println(ans)
}
