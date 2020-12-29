package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)

	for i := 0; i*2 <= len(name); i++ {
		if name[i] != name[len(name)-i-1] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
