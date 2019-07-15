package main

import (
	"fmt"
	"strconv"
)

func base(x, b int) string {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	ans := ""
	for x != 0 {
		r := abs(x % b)
		ans = strconv.Itoa(r) + ans
		x -= r
		x /= b
	}
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(base(x, -2))
}
