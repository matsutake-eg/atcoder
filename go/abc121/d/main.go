package main

import "fmt"

func xor(x int) (ans int) {
	switch x % 4 {
	case 0:
		ans = x
	case 1:
		ans = 1
	case 2:
		ans = 1 ^ x
	case 3:
		ans = 0
	}
	return
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(xor(a-1) ^ xor(b))
}
