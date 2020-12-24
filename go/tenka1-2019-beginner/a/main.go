package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a < c && c < b) || (b < c && c < a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
