package main

import "fmt"

func main() {
	var h1, w1, h2, w2 int
	fmt.Scan(&h1, &w1, &h2, &w2)

	if h1 == h2 || h1 == w2 || w1 == h2 || w1 == w2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
