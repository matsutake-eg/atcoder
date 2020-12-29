package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if l, r := a+b, c+d; l > r {
		fmt.Println("Left")
	} else if l < r {
		fmt.Println("Right")
	} else {
		fmt.Println("Balanced")
	}
}
