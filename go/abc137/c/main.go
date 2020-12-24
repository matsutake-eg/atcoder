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
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	n := scanInt()
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		s := scanString()
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		m[string(r)]++
	}

	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)
}
