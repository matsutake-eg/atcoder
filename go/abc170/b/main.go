package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if y > 4*x || y < x*2 || y%2 == 1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
