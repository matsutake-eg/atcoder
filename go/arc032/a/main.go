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

	sum := (1 + n) * n / 2
	if isPrime(sum) {
		fmt.Println("WANWAN")
	} else {
		fmt.Println("BOWWOW")
	}
}
