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
	p := make([]int, n)
	for i := range p {
		p[i] = scanInt()
	}
	sort.Ints(p)

	ans := 0
	for i := 0; i < k; i++ {
		ans += p[i]
	}
	fmt.Println(ans)
}
