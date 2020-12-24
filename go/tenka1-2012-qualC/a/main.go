package main

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	for i := 1; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	fmt.Println(ans)
}
