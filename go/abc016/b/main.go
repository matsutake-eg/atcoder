package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b == c && a-b == c {
		fmt.Println("?")
		return
	}

	if a+b == c {
		fmt.Println("+")
	} else if a-b == c {
		fmt.Println("-")
	} else {
		fmt.Println("!")
	}
}
