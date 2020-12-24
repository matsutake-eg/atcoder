package main

import "fmt"

func main() {
	var a, k int
	fmt.Scan(&a, &k)

	const limit = 2000000000000
	if k == 0 {
		fmt.Println(limit - a)
		return
	}

	ans := 0
	for {
		a += 1 + k*a
		ans++
		if 1 <= a/limit {
			fmt.Println(ans)
			return
		}
	}
}
