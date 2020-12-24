package main

import "fmt"

const MAX = 10000

func main() {
	var N, T int
	fmt.Scan(&N, &T)

	minCost := MAX
	var c, t int
	for i := 0; i < N; i++ {
		fmt.Scan(&c, &t)
		if t <= T && c < minCost {
			minCost = c
		}
	}

	if minCost == MAX {
		fmt.Println("TLE")
	} else {
		fmt.Println(minCost)
	}
}
