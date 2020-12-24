package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	ans := 1
	for i := 2; i < x; i++ {
		if i*i > x {
			break
		}
		for j := i * i; j <= x; j *= i {
			if j > ans {
				ans = j
			}
		}
	}
	fmt.Println(ans)
}
