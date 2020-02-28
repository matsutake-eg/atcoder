package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if v := a + b; v >= 10 {
		fmt.Println("error")
	} else {
		fmt.Println(v)
	}
}
