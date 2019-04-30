package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var ans int
	for i := 0; i <= c; i++ {
		if a*i <= b {
			ans = i
		}
	}
	fmt.Println(ans)
}
