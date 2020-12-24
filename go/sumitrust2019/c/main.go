package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x%100 > (x/100)*5 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
