package main

import "fmt"

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	if k*500 >= x {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
