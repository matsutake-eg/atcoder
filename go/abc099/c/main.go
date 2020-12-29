package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := n
	for i := 0; i <= n; i++ {
		sum := 0
		for x := i; x > 0; x /= 6 {
			sum += x % 6
		}
		for x := n - i; x > 0; x /= 9 {
			sum += x % 9
		}
		if sum < ans {
			ans = sum
		}
	}
	fmt.Println(ans)
}
