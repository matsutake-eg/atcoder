package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 7:
		fmt.Println("YES")
	case 5:
		fmt.Println("YES")
	case 3:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
