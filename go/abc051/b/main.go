package main

import "fmt"

func main() {
	var k, s int
	fmt.Scan(&k, &s)

	ans := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - x - y
			if z >= 0 && z <= k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
