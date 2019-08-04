package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i <= n; i++ {
		if i >= 1 && i <= 9 {
			ans++
		}
		if i >= 100 && i <= 999 {
			ans++
		}
		if i >= 10000 && i <= 99999 {
			ans++
		}
	}
	fmt.Println(ans)
}
