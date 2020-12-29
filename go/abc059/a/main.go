package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	fmt.Println(strings.ToUpper(a[:1] + b[:1] + c[:1]))
}
