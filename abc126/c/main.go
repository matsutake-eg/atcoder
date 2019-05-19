package main

import (
	"fmt"
	"math"
)

func log(x, max int) float64 {
	ans := 0
	for {
		if x >= max {
			return float64(ans)
		}
		x *= 2
		ans++
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var sum float64
	if n <= k {
		for i := 1; i <= n; i++ {
			sum += math.Pow(0.5, log(i, k))
		}
	} else {
		for i := 1; i <= k; i++ {
			sum += math.Pow(0.5, log(i, k))
		}
		sum += float64(n - k)
	}
	fmt.Println((1 / float64(n)) * sum)
}
