package main

import (
	"fmt"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)

	s := strconv.Itoa(n)
	c, _ := strconv.Atoi(string(s[0]))
	for i := 1; i < len(s); i++ {
		if s[i] != '9' {
			fmt.Println(c - 1 + (len(s)-1)*9)
			return
		}
	}
	fmt.Println(c + (len(s)-1)*9)
}
