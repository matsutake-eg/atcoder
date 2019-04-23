package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	is_exist := false
	for i := 1; i <= b; i++ {
		if i*a%b == c {
			is_exist = true
			break
		}
	}

	if is_exist {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
