package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	if a[len(a)-1] == b[0] && b[len(b)-1] == c[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
