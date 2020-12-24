package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if len(s) >= 4 && s[:4] == "YAKI" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
