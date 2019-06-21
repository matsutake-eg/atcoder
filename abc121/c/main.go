package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type shop struct {
	price int
	stock int
}

type shops []shop

func (s shops) Len() int {
	return len(s)
}

func (s shops) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s shops) Less(i, j int) bool {
	return s[i].price < s[j].price
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var s shops = make([]shop, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i].price, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		s[i].stock, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(s)

	count := 0
	sum := 0
	for i := 0; i < n; i++ {
		if v := count + s[i].stock; v >= m {
			sum += s[i].price * (m - count)
			break
		} else {
			count = v
			sum += s[i].price * s[i].stock
		}
	}
	fmt.Println(sum)
}
