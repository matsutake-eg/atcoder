package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	n, k := scanInt(), scanInt()
	g := scanInt()
	mx := g
	for i := 0; i < n-1; i++ {
		a := scanInt()
		mx = max(mx, a)
		g = gcd(g, a)
	}

	if k <= mx && k%g == 0 {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
