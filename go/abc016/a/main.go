package main

import "fmt"

func main() {
	var m, d int
	fmt.Scan(&m, &d)

	if m%d == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
