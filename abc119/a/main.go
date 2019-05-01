package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s <= "2019/04/30" {
		fmt.Println("Heisei")
	} else {
		fmt.Println("TBD")
	}
}
