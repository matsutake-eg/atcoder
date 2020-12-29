package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	switch c {
	case "O", "P", "K", "L":
		fmt.Println("Right")
	default:
		fmt.Println("Left")
	}
}
