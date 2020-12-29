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
	var x int
	fmt.Scan(&x)

	for {
		if isPrime(x) {
			fmt.Println(x)
			return
		}
		x++
	}
}
