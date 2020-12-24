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

	bs := ""
	for x != 0 {
		r := abs(x % b)
		bs = strconv.Itoa(r) + bs
		x -= r
		x /= b
	}
	if bs == "" {
		bs = "0"
	}
	return bs
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	fmt.Println(len(base(n, k)))
}
