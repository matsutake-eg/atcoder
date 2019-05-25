package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var (
		l, r int
		lMax = 0
		rMin = 100001
	)
	for i := 0; i < m; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		if l > lMax {
			lMax = l
		}
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		if r < rMin {
			rMin = r
		}
	}

	v := rMin - lMax + 1
	if v < 0 {
		v = 0
	}
	if v < n {
		fmt.Println(v)
	} else {
		fmt.Println(n)
	}
}
