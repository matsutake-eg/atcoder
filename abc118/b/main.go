package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	xs := make([]int, m+1)
	var k, a int
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&a)
			xs[a]++
		}
	}

	cnt := 0
	for _, v := range xs {
		if v == n {
			cnt++
		}
	}
	fmt.Println(cnt)
}
