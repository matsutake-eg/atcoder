package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b < c {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
