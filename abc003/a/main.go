package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0.0
	for i := 1; i <= n; i++ {
		ans += 10000 * float64(i) / float64(n)
	}
	fmt.Println(ans)
}
