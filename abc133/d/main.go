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
	a := make([]int, n)
	x := 0
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
		x += a[i]
	}

	for i := 2; i < n; i += 2 {
		x -= 2 * a[i-1]
	}
	fmt.Print(x)
	for i := 0; i < n-1; i++ {
		x = a[i]*2 - x
		fmt.Printf(" %v", x)
	}
	fmt.Println()
}
