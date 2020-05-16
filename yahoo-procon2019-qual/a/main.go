package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if 1+(k-1)*2 <= n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
