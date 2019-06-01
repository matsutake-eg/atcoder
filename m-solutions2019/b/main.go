package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	wins := strings.Count(s, "o")
	if 15-len(s)+wins >= 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
