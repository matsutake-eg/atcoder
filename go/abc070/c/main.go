package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x / gcd(x, y) * y
}

func main() {
	n := readInt()
	ans := 1
	for i := 0; i < n; i++ {
		t := readInt()
		ans = lcm(ans, t)
	}
	fmt.Println(ans)
}
