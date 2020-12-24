package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i*i < n; i++ {
		if x := n % i; x == 0 {
			if y := (n / i) - 1; y > i {
				ans += y
			}
		}
	}
	fmt.Println(ans)
}
