package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	for i := 1; i*i*i <= a; i++ {
		if i*i*i == a {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
