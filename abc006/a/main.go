package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n%3 == 0 || strings.ContainsRune(strconv.Itoa(n), '3') {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
