package main

import "fmt"

func main() {
	var l1, l2, l3 int
	fmt.Scan(&l1, &l2, &l3)

	if l1 == l2 && l2 == l3 {
		fmt.Println(l1)
	} else if l1 == l2 {
		fmt.Println(l3)
	} else if l2 == l3 {
		fmt.Println(l1)
	} else {
		fmt.Println(l2)
	}
}
