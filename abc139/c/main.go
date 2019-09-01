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
	ans := 0
	cnt := 0
	bh := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())
		if h <= bh {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > ans {
			ans = cnt
		}
		bh = h
	}
	fmt.Println(ans)
}
