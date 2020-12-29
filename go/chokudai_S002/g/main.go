package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return gcd(b, r)
}

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var a, b int
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())
		fmt.Println(gcd(a, b))
	}
}
