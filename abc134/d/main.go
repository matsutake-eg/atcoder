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
	a := make([]int, n)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	b := make([]int, n)
	ans := make([]int, 0, n)
	for i := len(a); i >= 1; i-- {
		x := 0
		for j := 2; j <= n; j++ {
			if i*j > n {
				break
			}
			x ^= b[i*j-1]
		}
		b[i-1] = x ^ a[i-1]
		if b[i-1] == 1 {
			ans = append(ans, i)
		}
	}

	fmt.Println(len(ans))

	wr := bufio.NewWriter(os.Stdout)
	for i, v := range ans {
		if i == len(ans)-1 {
			wr.WriteString(strconv.Itoa(v) + "\n")
			continue
		}
		wr.WriteString(strconv.Itoa(v) + " ")
	}
	wr.Flush()
}
