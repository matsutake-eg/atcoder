package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n int
		k int64
	)
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, n)
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
	}

	sum := int64(0)
	cnt := int64(0)
	l := -1
	for i := range as {
		sum += int64(as[i])
		if sum >= k {
			cnt += int64(len(as) - i)
			for l <= i {
				l++
				sum -= int64(as[l])
				if sum < k {
					break
				}
				cnt += int64(len(as) - i)
			}
		}
	}
	fmt.Println(cnt)
}
