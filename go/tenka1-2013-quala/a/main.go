package main

import "fmt"

func main() {
	x := 42
	for x < 130000000 {
		x += x
	}
	fmt.Println(x)
}
