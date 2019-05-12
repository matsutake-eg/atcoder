package main

import "fmt"

func main() {
	var R, G, B, N int
	fmt.Scan(&R, &G, &B, &N)

	count := 0
	for r := 0; r*R <= N; r++ {
		for g := 0; r*R+g*G <= N; g++ {
			if (N-r*R-g*G)%B == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
