package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		ans += i
	}
	fmt.Println(ans)
}
