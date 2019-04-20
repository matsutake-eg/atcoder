package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < c && c < b {
		fmt.Println("Yes")
	} else if b < c && c < a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
