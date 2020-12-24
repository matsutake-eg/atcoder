package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	fmt.Println(string(c[0] + 1))
}
