package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	x, _ := strconv.Atoi(a + b)
	fmt.Println(x * 2)
}
