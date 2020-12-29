package main

import (
	"fmt"
	"strconv"
	"strings"
)

var N int

func dfs(x int) int {
	if x > N {
		return 0
	}
	cnt := 0
	if contains(strconv.Itoa(x), "753") {
		cnt++
	}
	cnt += dfs(x*10 + 7)
	cnt += dfs(x*10 + 5)
	cnt += dfs(x*10 + 3)
	return cnt
}
func contains(t, s string) bool {
	for _, r := range s {
		if !strings.ContainsRune(t, r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scan(&N)

	fmt.Println(dfs(0))
}
