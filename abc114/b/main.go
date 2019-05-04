package main

import (
	"fmt"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var s string
	fmt.Scan(&s)

	var x int
	min := 1000
	for i := 0; i+2 < len(s); i++ {
		x, _ = strconv.Atoi(s[i : i+3])
		if v := abs(x - 753); v < min {
			min = v
		}
	}
	fmt.Println(min)
}
