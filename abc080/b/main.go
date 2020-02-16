package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0
	t := n
	for t > 0 {
		s += t % 10
		t /= 10
	}
	if n%s == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
