package main

import (
	"fmt"
	"strconv"
)

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

	if isPrime(n) {
		fmt.Println("Prime")
	} else if n == 1 {
		fmt.Println("Not Prime")
	} else {
		sn := strconv.Itoa(n)
		switch sn[len(sn)-1] {
		case '1', '3', '7', '9':
			sum := 0
			for _, r := range sn {
				x, _ := strconv.Atoi(string(r))
				sum += x
			}
			if sum%3 != 0 {
				fmt.Println("Prime")
				return
			}
		}
		fmt.Println("Not Prime")
	}
}
