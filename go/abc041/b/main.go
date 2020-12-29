package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	const mod = 1000000007
	ans := a * b % mod
	ans = ans * c % mod
	fmt.Println(ans)
}
