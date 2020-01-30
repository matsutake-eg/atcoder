package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	x, _ := strconv.Atoi(a + b)
	for i := 1; i*i <= x; i++ {
		if i*i == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
