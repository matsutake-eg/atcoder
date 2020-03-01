package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	i := 0
	for {
		if (a+i)%b == 0 {
			fmt.Println(i)
			return
		}
		i++
	}
}
