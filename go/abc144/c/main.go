package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := n - 1
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		if v := i - 1 + (n / i) - 1; v < ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
