package main

import "fmt"

func main() {
	var (
		a, t float32
		b    int
	)
	fmt.Scan(&a, &b, &t)

	sum := 0
	for i := 1; a*float32(i) < t+0.5; i++ {
		sum += b
	}
	fmt.Println(sum)
}
