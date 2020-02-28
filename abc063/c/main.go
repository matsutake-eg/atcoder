package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	ans := 0
	mn := math.MaxInt64
	for i := 0; i < n; i++ {
		s := scanInt()
		ans += s
		if s%10 != 0 {
			mn = min(mn, s)
		}
	}

	if ans%10 != 0 {
		fmt.Println(ans)
	} else if mn == math.MaxInt64 {
		fmt.Println(0)
	} else {
		fmt.Println(ans - mn)
	}
}
