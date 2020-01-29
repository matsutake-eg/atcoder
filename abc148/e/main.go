package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 1 {
		fmt.Println(0)
		return
	}

	ans := 0
	for i := 10; i <= n; i *= 5 {
		ans += n / i
	}
	fmt.Println(ans)
}
