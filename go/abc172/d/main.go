package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := (n + 1) * n / 2
	for i := 2; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			ans += i * j
		}
	}
	fmt.Println(ans)
}
