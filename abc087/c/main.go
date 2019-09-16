package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a1 := make([]int, n)
	for i := range a1 {
		fmt.Scan(&a1[i])
		if i > 0 {
			a1[i] += a1[i-1]
		}
	}
	a2 := make([]int, n)
	for i := range a2 {
		fmt.Scan(&a2[i])
	}
	for i := len(a2) - 1; i >= 0; i-- {
		if i < len(a2)-1 {
			a2[i] += a2[i+1]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if v := a1[i] + a2[i]; v > ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
