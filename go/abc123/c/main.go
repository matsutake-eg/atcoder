package main

import (
	"fmt"
	"math"
)

func main() {
	var n, min int
	fmt.Scan(&n, &min)
	xs := make([]int, 4)
	for i := range xs {
		fmt.Scan(&xs[i])
		if xs[i] < min {
			min = xs[i]
		}
	}
	fmt.Println(int(math.Ceil(float64(n)/float64(min))) + 4)
}
