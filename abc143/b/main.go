package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	d := make([]int, n)
	for i := range d {
		fmt.Scan(&d[i])
	}

	ans := 0
	for i := range d {
		for j := i + 1; j < len(d); j++ {
			ans += d[i] * d[j]
		}
	}
	fmt.Println(ans)
}
