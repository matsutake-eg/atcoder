package main

import "fmt"

func main() {
	var a1, a2, a3 int
	fmt.Scan(&a1, &a2, &a3)

	if a1+a2+a3 >= 22 {
		fmt.Println("bust")
		return
	}
	fmt.Println("win")
}
