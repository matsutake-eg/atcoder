package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if n == i*j {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
