package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vs[i])
	}
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cs[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if v := vs[i] - cs[i]; v > 0 {
			sum += v
		}
	}
	fmt.Println(sum)
}
