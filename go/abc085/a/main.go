package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println("2018" + s[4:])
}
