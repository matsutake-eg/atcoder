package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	x := n / k
	ans := x * x * x
	if k%2 == 0 {
		x = n/(k/2) - x
		ans += x * x * x
	}
	fmt.Println(ans)
}
