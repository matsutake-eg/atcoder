package main

import (
	"fmt"
	"math"
)

func main() {
	var n, min, x int
	fmt.Scan(&n, &min)
	for i := 1; i < 5; i++ {
		fmt.Scan(&x)
		if x < min {
			min = x
		}
	}
	fmt.Println(int(math.Ceil(float64(n)/float64(min))) + 4)
}
