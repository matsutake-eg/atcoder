package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if (a == "H" && b == "H") || (a == "D" && b == "D") {
		fmt.Println("H")
	} else {
		fmt.Println("D")
	}
}
