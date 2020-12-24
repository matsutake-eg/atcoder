package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if i*i*i*i == x {
			fmt.Println(i)
			return
		}
	}
}
