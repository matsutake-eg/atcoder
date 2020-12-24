package main

import "fmt"

func main() {
	var c1, c2, c3 string
	fmt.Scan(&c1, &c2, &c3)

	fmt.Println(string(c1[0]) + string(c2[1]) + string(c3[2]))
}
