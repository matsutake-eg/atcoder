package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var x, y, z, k int
	fmt.Scan(&x, &y, &z, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, x)
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(as)))
	bs := make([]int, y)
	for i := range bs {
		sc.Scan()
		bs[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bs)))
	cs := make([]int, z)
	for i := range cs {
		sc.Scan()
		cs[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cs)))

	ans := make([]int, 0, x+y+z)
	for ia, a := range as {
		for ib, b := range bs {
			for ic, c := range cs {
				if (ia+1)*(ib+1)*(ic+1) > k {
					break
				}
				ans = append(ans, a+b+c)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ans)))
	for _, v := range ans[:k] {
		fmt.Println(v)
	}
}
