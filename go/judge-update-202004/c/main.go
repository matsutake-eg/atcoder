package main

import "fmt"

var a, b, c int

func dfs(x, y, z int) int {
	if x == a && y == b && z == c {
		return 1
	}

	ans := 0
	if x < a {
		ans += dfs(x+1, y, z)
	}
	if y < x && y < b {
		ans += dfs(x, y+1, z)
	}
	if z < y && z < c {
		ans += dfs(x, y, z+1)
	}
	return ans
}

func main() {
	fmt.Scan(&a, &b, &c)

	fmt.Println(dfs(0, 0, 0))
}
