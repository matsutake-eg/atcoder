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

type number struct{ a, b int }
type numbers []number

func (n numbers) Len() int           { return len(n) }
func (n numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n numbers) Less(i, j int) bool { return n[i].a < n[j].a }

func main() {
	n, k := scanInt(), scanInt()
	nums := numbers(make([]number, n))
	for i := range nums {
		nums[i].a, nums[i].b = scanInt(), scanInt()
	}

	sort.Sort(nums)
	cnt := 0
	for _, num := range nums {
		if cnt+num.b >= k {
			fmt.Println(num.a)
			return
		}
		cnt += num.b
	}
}
