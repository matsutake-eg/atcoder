package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type arm struct{ l, r int }
type arms []arm

func (a arms) Len() int           { return len(a) }
func (a arms) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a arms) Less(i, j int) bool { return a[i].r < a[j].r }

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
	n := scanInt()
	as := arms(make([]arm, n))
	for i := range as {
		x, l := scanInt(), scanInt()
		as[i].l = x - l
		as[i].r = x + l

	}
	sort.Sort(as)

	ans := 1
	r := as[0].r
	for i := 1; i < n; i++ {
		if as[i].l >= r {
			ans++
			r = as[i].r
		}
	}
	fmt.Println(ans)
}
