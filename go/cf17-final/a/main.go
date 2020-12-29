package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	switch s {
	case "AKIHABARA", "KIHABARA", "AKIHBARA", "AKIHABRA", "AKIHABAR", "KIHBARA", "KIHABRA", "KIHABAR", "AKIHBRA", "AKIHBAR", "AKIHABR", "KIHBRA", "KIHBAR", "KIHABR", "AKIHBR", "KIHBR":
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
