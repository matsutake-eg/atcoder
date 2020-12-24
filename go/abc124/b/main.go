package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	max := 0
	ans := 0
	h := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		if h >= max {
			max = h
			ans++
		}
	}
	fmt.Println(ans)
}
