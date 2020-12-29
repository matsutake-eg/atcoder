package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var ra, rb, rc int
	if a >= b && a >= c {
		ra = 1
	} else if a < b && a < c {
		ra = 3
	} else {
		ra = 2
	}
	if b >= a && b >= c {
		rb = 1
	} else if b < a && b < c {
		rb = 3
	} else {
		rb = 2
	}
	if c >= a && c >= b {
		rc = 1
	} else if c < a && c < b {
		rc = 3
	} else {
		rc = 2
	}
	fmt.Println(ra)
	fmt.Println(rb)
	fmt.Println(rc)
}
