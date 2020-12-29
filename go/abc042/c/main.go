package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var n int
var nums []string

func dfs(s string) int {
	if len(s) > len(strconv.Itoa(n))+1 {
		return math.MaxInt64
	}
	sv, _ := strconv.Atoi(s)
	if sv >= n {
		return sv
	}

	ans := math.MaxInt64
	for _, v := range nums {
		ans = min(ans, dfs(s+v))
	}
	return ans
}

func main() {
	n = scanInt()
	k := scanInt()
	d := make(map[int]bool)
	for i := 0; i < k; i++ {
		x := scanInt()
		d[x] = true
	}

	nums = make([]string, 0, 10)
	for i := 0; i <= 9; i++ {
		if !d[i] {
			nums = append(nums, strconv.Itoa(i))
		}
	}
	fmt.Println(dfs(""))
}
