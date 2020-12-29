package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	ans := float64(0)
	for i := 0; i < n; i++ {
		sum := float64(0)
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			sum += math.Sqrt(math.Pow(x[i]-x[j], 2) + math.Pow(y[i]-y[j], 2))
		}
		ans += sum / float64(n)
	}
	fmt.Println(ans)
}
