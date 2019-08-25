package main

import "fmt"

func main() {
	var n, d, x int
	fmt.Scan(&n, &d, &x)

	ans := x
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		for i := 1; i <= d; i += a {
			ans++
		}
	}
	fmt.Println(ans)
}
