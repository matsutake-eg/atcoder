package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	loses := strings.Count(s, "x")
	if 15-loses >= 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
