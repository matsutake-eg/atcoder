package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x%y == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(x)
	}
}
