package main

import (
	"fmt"
	"strings"
)

var (
	n int
	m []map[string]int
)

func dfs(cnt int, last3 string) int {
	if v, ok := m[cnt][last3]; ok {
		return v
	}

	if cnt == n {
		return 1
	}

	ans := 0
	for _, r := range "ACGT" {
		if isNotAGC(last3 + string(r)) {
			ans = (ans + dfs(cnt+1, last3[1:]+string(r))) % 1000000007
		}
	}

	m[cnt][last3] = ans
	return ans
}

func isNotAGC(last4 string) bool {
	s := make([]rune, 4)
	for i := range last4 {
		copy(s, []rune(last4))
		if i < 3 {
			s[i], s[i+1] = s[i+1], s[i]
		}

		if strings.Contains(string(s), "AGC") {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scan(&n)

	m = make([]map[string]int, n+1)
	for i := range m {
		m[i] = make(map[string]int)
	}
	fmt.Println(dfs(0, "TTT"))
}
