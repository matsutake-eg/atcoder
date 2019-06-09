package main

import "fmt"

func main() {
	var (
		xs       = make([]int, 3)
		max, sum int
	)
	for i := range xs {
		fmt.Scan(&xs[i])
		sum += xs[i]
		if xs[i] > max {
			max = xs[i]
		}
	}
	fmt.Println(sum - max)
}
