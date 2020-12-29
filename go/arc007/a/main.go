package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, s string
	fmt.Scan(&x, &s)

	fmt.Println(strings.ReplaceAll(s, x, ""))
}
