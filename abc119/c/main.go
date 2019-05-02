package main

import "fmt"

const INF = 1000000000

var (
	N, A, B, C int
	L          []int
)

func dfs(cur, a, b, c int) int {
	if cur == N {
		if a != 0 && b != 0 && c != 0 {
			return abs(A-a) + abs(B-b) + abs(C-c) - 30
		}
		return INF
	}
	n1 := dfs(cur+1, a, b, c)
	n2 := dfs(cur+1, a+L[cur], b, c) + 10
	n3 := dfs(cur+1, a, b+L[cur], c) + 10
	n4 := dfs(cur+1, a, b, c+L[cur]) + 10
	return min(n1, n2, n3, n4)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func main() {
	fmt.Scan(&N, &A, &B, &C)
	L = make([]int, N)
	for i := range L {
		fmt.Scan(&L[i])
	}

	fmt.Println(dfs(0, 0, 0, 0))
}
