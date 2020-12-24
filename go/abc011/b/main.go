package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(strings.ToUpper(s[:1]) + strings.ToLower(s[1:]))
}
