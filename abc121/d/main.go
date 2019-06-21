package main

import "fmt"

func xor(x int64) (ans int64) {
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
	var a, b int64
	fmt.Scan(&a, &b)

	fmt.Println(xor(a-1) ^ xor(b))
}
