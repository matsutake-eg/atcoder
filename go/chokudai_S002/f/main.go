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
	var a, b int
	xm := make(map[string]int)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())
		if a < b {
			xm[strconv.Itoa(a)+strconv.Itoa(b)] = 0
		} else {
			xm[strconv.Itoa(b)+strconv.Itoa(a)] = 0
		}
	}
	fmt.Println(len(xm))
}
