package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	if len(a) > 1 {
		fmt.Println(a[:len(a)-1])
	} else if a == "a" {
		fmt.Println(-1)
	} else {
		fmt.Println("a")
	}
}
