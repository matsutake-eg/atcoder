package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	sc := bufio.NewScanner(os.Stdin)
	ans := make([]int, n)
	for i := 0; i < q; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		ans[a-1]++
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		if k-(q-v) > 0 {
			wr.WriteString("Yes\n")
			continue
		}
		wr.WriteString("No\n")
	}
	wr.Flush()
}
