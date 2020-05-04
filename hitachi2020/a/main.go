package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if len(s)%2 == 1 {
		fmt.Println("No")
		return
	}

	for i := 0; i+1 < len(s); i += 2 {
		if s[i:i+2] != "hi" {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
