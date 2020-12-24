package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 50000; i++ {
		if i*108/100 == n {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(":(")
}
