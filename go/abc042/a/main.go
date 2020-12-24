package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a == 7 && b == 5 && c == 5) || (a == 5 && b == 7 && c == 5) || (a == 5 && b == 5 && c == 7) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
