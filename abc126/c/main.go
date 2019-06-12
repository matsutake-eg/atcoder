package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	sum := 0.0
	limit := n
	if n > k {
		limit = k
		sum += float64(n - k)
	}
	for i := 1; i <= limit; i++ {
		sum += math.Pow(0.5, math.Ceil(math.Log2(float64(k)/float64(i))))
	}
	fmt.Println((1 / float64(n)) * sum)
}
