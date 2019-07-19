package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sn := 0
	for _, r := range strconv.Itoa(n) {
		x, _ := strconv.Atoi(string(r))
		sn += x
	}

	if n%sn == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
