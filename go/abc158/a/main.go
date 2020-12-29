package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s == "AAA" || s == "BBB" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
