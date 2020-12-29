package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	sum := 0
	max := 99
	for i := range p {
		fmt.Scan(&p[i])
		sum += p[i]
		if p[i] > max {
			max = p[i]
		}
	}
	fmt.Println(sum - max/2)
}
