package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if b < c || d < a {
		fmt.Println(0)
		return
	}

	if a < c {
		if b < d {
			fmt.Println(b - c)
		} else {
			fmt.Println(d - c)
		}
	} else {
		if b < d {
			fmt.Println(b - a)
		} else {
			fmt.Println(d - a)
		}
	}
}
