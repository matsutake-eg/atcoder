package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	fmt.Println("A" + b[:1] + "C")
}
