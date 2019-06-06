package main

import "fmt"

func main() {
	xs := make([]int, 5)
	var sum, max int
	for i := range xs {
		fmt.Scan(&xs[i])
		if xs[i]%10 == 0 {
			sum += xs[i]
		} else {
			sum += (xs[i]/10 + 1) * 10
			if v := 10 - xs[i]%10; v > max {
				max = v
			}
		}
	}
	fmt.Println(sum - max)
}
