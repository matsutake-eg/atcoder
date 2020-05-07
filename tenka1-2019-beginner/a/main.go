package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (c > a && c < b) || (c > b && c < a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
