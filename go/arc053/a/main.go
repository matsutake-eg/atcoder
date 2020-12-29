package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i+1 < h {
				ans++
			}
			if j+1 < w {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
