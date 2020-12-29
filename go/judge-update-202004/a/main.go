package main

import "fmt"

func main() {
	var s, l, r int
	fmt.Scan(&s, &l, &r)

	if s >= l && s <= r {
		fmt.Println(s)
	} else if s < l {
		fmt.Println(l)
	} else {
		fmt.Println(r)
	}
}
