package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, k := scanInt(), scanInt()
	type number struct{ a, b int }
	nums := make([]number, n)
	for i := range nums {
		nums[i].a, nums[i].b = scanInt(), scanInt()
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i].a < nums[j].a })
	cnt := 0
	for _, num := range nums {
		if cnt+num.b >= k {
			fmt.Println(num.a)
			return
		}
		cnt += num.b
	}
}
