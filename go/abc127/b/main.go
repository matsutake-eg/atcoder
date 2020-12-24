package main

import "fmt"

func main() {
	var r, D, x int
	fmt.Scan(&r, &D, &x)

	for i := 0; i < 10; i++ {
		x = r*x - D
		fmt.Println(x)
	}
}
