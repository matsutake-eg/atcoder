package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	sc := bufio.NewScanner(os.Stdin)
	min := 10000
	for i := 0; i < n; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		x -= m
		if m < min {
			min = m
		}
	}
	fmt.Println(n + x/min)
}
