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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	ans := 1000000000
	for i := 0; i < n; i++ {
		a := readInt()
		cnt := 0
		for a > 0 {
			if a%2 == 1 {
				break
			}
			a /= 2
			cnt++
		}
		ans = min(ans, cnt)
	}
	fmt.Println(ans)
}
