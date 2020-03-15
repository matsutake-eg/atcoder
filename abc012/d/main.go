package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = 1001 * 300
		}
	}

	m := scanInt()
	for i := 0; i < m; i++ {
		a, b, t := scanInt(), scanInt(), scanInt()
		a--
		b--
		dist[a][b] = t
		dist[b][a] = t
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	ans := math.MaxInt64
	for i := range dist {
		x := 0
		for j := range dist[i] {
			if i == j {
				continue
			}

			x = max(x, dist[i][j])
		}
		ans = min(ans, x)
	}
	fmt.Println(ans)
}
