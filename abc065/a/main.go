package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	if v := b - a; v <= 0 {
		fmt.Println("delicious")
	} else if v <= x {
		fmt.Println("safe")
	} else {
		fmt.Println("dangerous")
	}
}
