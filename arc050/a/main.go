package main

import (
	"fmt"
	"strings"
)

func main() {
	var C, c string
	fmt.Scan(&C, &c)

	if strings.ToLower(C) == c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
