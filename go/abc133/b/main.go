package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, d)
		for j := range x[i] {
			fmt.Scan(&x[i][j])
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := 0
			for k := 0; k < d; k++ {
				sum += (x[i][k] - x[j][k]) * (x[i][k] - x[j][k])
			}
			for k := 0; k <= sum; k++ {
				if v := k * k; v == sum {
					cnt++
				} else if v > sum {
					break
				}
			}
		}
	}
	fmt.Println(cnt)
}
