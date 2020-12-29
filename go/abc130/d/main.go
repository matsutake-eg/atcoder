package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, n)
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
	}

	sum := 0
	cnt := 0
	l := -1
	for i := range as {
		sum += as[i]
		if sum >= k {
			cnt += len(as) - i
			for l <= i {
				l++
				sum -= as[l]
				if sum < k {
					break
				}
				cnt += len(as) - i
			}
		}
	}
	fmt.Println(cnt)
}
