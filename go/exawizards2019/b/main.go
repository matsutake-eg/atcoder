package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	r := strings.Count(s, "R")
	b := n - r
	if r > b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
