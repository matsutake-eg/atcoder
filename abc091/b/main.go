package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	mp := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		mp[s]++
	}

	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var t string
		fmt.Scan(&t)
		mp[t]--
	}

	ans := 0
	for _, v := range mp {
		if v > ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
