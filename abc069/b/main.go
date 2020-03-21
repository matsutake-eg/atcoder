package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(s[:1] + strconv.Itoa(len(s)-2) + s[len(s)-1:])
}
