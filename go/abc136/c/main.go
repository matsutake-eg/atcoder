package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	bh := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())
		if bh > h {
			fmt.Println("No")
			return
		}
		if h == bh {
			bh = h
		} else {
			bh = h - 1
		}
	}
	fmt.Println("Yes")
}
