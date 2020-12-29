package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)

	if y%400 == 0 {
		fmt.Println("YES")
		return
	}
	if y%100 == 0 {
		fmt.Println("NO")
		return
	}
	if y%4 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
