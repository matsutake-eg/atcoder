package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 1
	for i := 2; i*i <= n; i++ {
		if i*i == n {
			sum += i
		} else if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	if n == 1 {
		fmt.Println("Deficient")
	} else if sum == n {
		fmt.Println("Perfect")
	} else if sum < n {
		fmt.Println("Deficient")
	} else {
		fmt.Println("Abundant")
	}
}
