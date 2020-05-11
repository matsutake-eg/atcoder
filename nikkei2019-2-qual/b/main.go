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

func main() {
	n := scanInt()
	d0 := scanInt()
	if d0 > 0 {
		fmt.Println(0)
		return
	}

	xm := make(map[int]int, n)
	for i := 1; i < n; i++ {
		d := scanInt()
		xm[d]++
	}

	const mod = 998244353
	ans := 1
	for k, v := range xm {
		if k-1 != 0 && xm[k-1] == 0 {
			fmt.Println(0)
			return
		}

		if xm[k-1] >= 2 {
			for i := 0; i < v; i++ {
				ans *= xm[k-1]
				ans %= mod
			}
		}
	}
	fmt.Println(ans)
}
