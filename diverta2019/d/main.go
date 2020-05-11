package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			if m := n/i - 1; m > i {
				ans += m
			}
		}
	}
	fmt.Println(ans)
}
