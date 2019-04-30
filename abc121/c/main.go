package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type shop struct {
	PRICE int
	STOCK int
}

type shops []shop

func (s shops) Len() int {
	return len(s)
}

func (s shops) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s shops) Less(i, j int) bool {
	return s[i].PRICE < s[j].PRICE
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var s shops = make([]shop, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i].PRICE, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		s[i].STOCK, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(s)

	count := 0
	sum := 0
	for i := 0; i < n; i++ {
		if v := count + s[i].STOCK; v >= m {
			sum += s[i].PRICE * (m - count)
			break
		} else {
			count = v
			sum += s[i].PRICE * s[i].STOCK
		}
	}
	fmt.Println(sum)
}
