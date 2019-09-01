package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b == 1 {
		fmt.Println(0)
		return
	}
	ans := 1
	t := a
	for t < b {
		t--
		t += a
		ans++
	}
	fmt.Println(ans)
}
