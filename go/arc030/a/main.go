package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k <= n/2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
