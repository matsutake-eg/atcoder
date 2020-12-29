package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 60 {
		fmt.Println("Bad")
	} else if n < 90 {
		fmt.Println("Good")
	} else if n < 100 {
		fmt.Println("Great")
	} else {
		fmt.Println("Perfect")
	}
}
