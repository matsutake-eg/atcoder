package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a <= 0 && b >= 0 {
		fmt.Println("Zero")
	} else if a > 0 || (b-a+1)%2 == 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}
