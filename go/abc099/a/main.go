package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 999 {
		fmt.Println("ABD")
		return
	}
	fmt.Println("ABC")
}
