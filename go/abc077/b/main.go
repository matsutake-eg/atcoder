package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for i := 1; i*i <= n; i++ {
		ans = i * i
	}
	fmt.Println(ans)
}
