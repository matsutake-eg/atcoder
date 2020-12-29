package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Print(strings.Count(s, "1"))
}
