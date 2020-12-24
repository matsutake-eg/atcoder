package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, ng1, ng2, ng3 int
	fmt.Scan(&n, &ng1, &ng2, &ng3)

	if n == ng1 || n == ng2 || n == ng3 {
		fmt.Println("NO")
		return
	}

	ans := make([]int, n+1)
	for i := range ans[:n] {
		ans[i] = math.MaxInt64
	}

	nums := []int{1, 2, 3}
	for i := n; i >= 0; i-- {
		if i == ng1 || i == ng2 || i == ng3 || ans[i] == math.MaxInt64 {
			continue
		}

		for _, num := range nums {
			if v := i - num; v >= 0 && v != ng1 && v != ng2 && v != ng3 {
				ans[v] = min(ans[v], ans[i]+1)
			}
		}
	}

	if ans[0] <= 100 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
