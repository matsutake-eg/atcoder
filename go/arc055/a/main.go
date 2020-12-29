package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("1" + strings.Repeat("0", n-1) + "7")
}
