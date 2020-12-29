package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)

	k++
	if b-a <= 2 {
		fmt.Println(k)
	} else {
		k -= a
		fmt.Println(a + k/2*(b-a) + k%2)
	}
}
