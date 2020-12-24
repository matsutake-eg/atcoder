package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	c0 := 0
	c1 := 0
	for _, v := range s {
		if v == '0' {
			c0++
		} else {
			c1++
		}
	}

	if c1 > c0 {
		fmt.Println(c0 * 2)
	} else {
		fmt.Println(c1 * 2)
	}
}
