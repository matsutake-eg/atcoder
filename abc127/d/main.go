package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type card struct{ number, stock int }
type cards []card

func (c cards) Len() int           { return len(c) }
func (c cards) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c cards) Less(i, j int) bool { return c[i].number > c[j].number }

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, n)
	sum := 0
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
		sum += as[i]
	}
	sort.Ints(as)

	cs := cards(make([]card, m))
	for i := range cs {
		sc.Scan()
		cs[i].stock, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		cs[i].number, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(cs)

	var p int
lb:
	for _, v := range cs {
		for i := 0; i < v.stock; i++ {
			if p >= len(as) || v.number < as[p] {
				break lb
			}
			sum += v.number - as[p]
			p++
		}
	}
	fmt.Println(sum)
}
