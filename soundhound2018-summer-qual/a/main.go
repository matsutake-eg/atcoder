package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a+b == 15 {
		fmt.Println("+")
	} else if a*b == 15 {
		fmt.Println("*")
	} else {
		fmt.Println("x")
	}
}
