package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var edge map[int]map[int]int
var ans []int

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(e, c, p int) {
	k := 1
	for to, idx := range edge[e] {
		if to == p {
			continue
		}
		if k == c {
			k++
		}
		ans[idx] = k
		k++
		dfs(to, ans[idx], e)
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	ans = make([]int, n-1)
	edge = make(map[int]map[int]int)

	for i := 0; i < n-1; i++ {
		a, b := readInt(), readInt()
		if _, ok := edge[a]; !ok {
			edge[a] = make(map[int]int)
		}
		edge[a][b] = i
		if _, ok := edge[b]; !ok {
			edge[b] = make(map[int]int)
		}
		edge[b][a] = i
	}

	dfs(1, -1, -1)
	mx := 0
	for _, v := range ans {
		mx = max(mx, v)
	}

	var wr = bufio.NewWriter(os.Stdout)
	wr.WriteString(strconv.Itoa(mx) + "\n")
	for i, v := range ans {
		wr.WriteString(strconv.Itoa(v))
		if i != len(ans)-1 {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
