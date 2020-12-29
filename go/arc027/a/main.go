package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	fmt.Println((18-h)*60 - m)
}
