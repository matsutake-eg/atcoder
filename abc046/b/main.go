package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ans := k
	for i := 0; i < n-1; i++ {
		ans *= k - 1
	}
	fmt.Println(ans)
}
