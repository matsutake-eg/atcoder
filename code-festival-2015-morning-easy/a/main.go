package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; ; i++ {
		if v := i * i; v >= n {
			fmt.Println(v - n)
			return
		}
	}
}
