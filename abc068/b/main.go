package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 1
	for ans*2 <= n {
		ans *= 2
	}
	fmt.Println(ans)
}
