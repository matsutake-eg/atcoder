package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans, x := 0, 1
	for x <= n {
		x *= 9
	}
	n -= x
	ans++
	if n == 0 {
		fmt.Println(ans)
		return
	}
}
