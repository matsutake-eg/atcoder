package main

import "fmt"

var (
	xs []int
	s  [][]int
	p  []int
)

func dfs(cur int, sw int) int {
	if cur != 0 {
		xs[cur-1] = sw
	}
	if cur == len(xs) {
		for i := range s {
			var light int
			for j := range s[i] {
				light += xs[s[i][j]-1]
			}
			if light%2 != p[i] {
				return 0
			}
		}
		return 1
	}
	sum := 0
	sum += dfs(cur+1, 0)
	sum += dfs(cur+1, 1)
	return sum
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s = make([][]int, m)
	var k int
	for i := range s {
		fmt.Scan(&k)
		s[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&s[i][j])
		}
	}
	p = make([]int, m)
	for i := range p {
		fmt.Scan(&p[i])
	}

	xs = make([]int, n)
	fmt.Println(dfs(0, 0))
}
