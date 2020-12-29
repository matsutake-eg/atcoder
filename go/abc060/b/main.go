package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	for i := 1; i < b; i++ {
		if a*i%b == c {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
