package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type sortRune []rune

func (r sortRune) Len() int           { return len(r) }
func (r sortRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r sortRune) Less(i, j int) bool { return r[i] < r[j] }

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		r := sortRune([]rune(sc.Text()))
		sort.Sort(r)
		m[string(r)]++
	}

	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)
}
