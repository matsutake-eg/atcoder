package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func max(x ...int) int {
	mv := x[0]
	for _, v := range x[1:] {
		if v > mv {
			mv = v
		}
	}
	return mv
}

func main() {
	n := readInt()
	m := make(map[int]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		a := readInt()
		m[a-1]++
		m[a]++
		m[a+1]++
		ans = max(ans, m[a-1], m[a], m[a+1])
	}
	fmt.Println(ans)
}
