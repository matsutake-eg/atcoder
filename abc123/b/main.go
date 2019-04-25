package main

import "fmt"

func main() {
	var x, p int
	sum := 0
	max := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&x)
		sum += x
		if v := x % 10; v != 0 {
			p = 10 - v
			sum += p
			if p > max {
				max = p
			}
		}
	}
	fmt.Println(sum - max)
}
