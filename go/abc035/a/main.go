package main

import "fmt"

func main() {
	var w, h int
	fmt.Scan(&w, &h)

	if w%16 == 0 && h%9 == 0 {
		fmt.Println("16:9")
	} else {
		fmt.Println("4:3")
	}
}
