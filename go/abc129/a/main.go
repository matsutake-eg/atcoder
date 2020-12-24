package main

import "fmt"

func main() {
	xs := make([]int, 3)
	sum := 0
	max := 0
	for i := range xs {
		fmt.Scan(&xs[i])
		sum += xs[i]
		if xs[i] > max {
			max = xs[i]
		}
	}
	fmt.Println(sum - max)
}
