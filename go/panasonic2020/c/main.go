package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a >= c || b >= c {
		fmt.Println("No")
		return
	}

	if 4*a*b < (c-a-b)*(c-a-b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
