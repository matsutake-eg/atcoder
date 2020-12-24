package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i < 10000; i++ {
		if int(float64(i)*0.08) == a && int(float64(i)*0.1) == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
