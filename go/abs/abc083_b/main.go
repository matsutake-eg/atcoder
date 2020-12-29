package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	total := 0
	var sum int

	for i := 1; i <= n; i++ {
		sum = 0
		for _, r := range strconv.Itoa(i) {
			sum += int(r - '0')
		}

		if sum >= a && sum <= b {
			total += i
		}
	}

	fmt.Println(total)
}
