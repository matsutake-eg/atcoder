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
	var (
		a, b int
		sum  = 0
	)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())
		if a > b {
			sum += a
		} else {
			sum += b
		}
	}
	fmt.Println(sum)
}
