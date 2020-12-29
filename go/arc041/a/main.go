package main

import "fmt"

func main() {
	var x, y, k int
	fmt.Scan(&x, &y, &k)

	if k <= y {
		fmt.Println(x + k)
	} else {
		fmt.Println(y + x - (k - y))
	}
}
