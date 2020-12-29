package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n/4; i++ {
		for j := 0; j <= n/7; j++ {
			if i*4+j*7 == n {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
