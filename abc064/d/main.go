package main

import (
	"fmt"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	c1, c2 := 0, 0
	for _, r := range s {
		if r == '(' {
			c1++
		} else {
			if c1 > 0 {
				c1--
			} else {
				c2++
			}
		}
	}

	fmt.Println(strings.Repeat("(", c2) + s + strings.Repeat(")", c1))
}
