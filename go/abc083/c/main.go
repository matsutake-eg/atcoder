package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	ans := 0
	for x <= y {
		x *= 2
		ans++
	}
	fmt.Println(ans)
}
