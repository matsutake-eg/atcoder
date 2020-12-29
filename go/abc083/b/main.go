package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := 0
	for i := 1; i <= n; i++ {
		t := i
		sum := 0
		for t > 0 {
			sum += t % 10
			t /= 10
		}
		if sum >= a && sum <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
