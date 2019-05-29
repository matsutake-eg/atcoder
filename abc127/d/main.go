package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Card struct{ Number, Stock int }
type Cards []Card

func (c Cards) Len() int           { return len(c) }
func (c Cards) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool { return c[i].Number > c[j].Number }

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, n)
	var sum int64
	for i := range as {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
		sum += int64(as[i])
	}
	sort.Ints(as)
	var cs Cards = make([]Card, m)
	for i := range cs {
		sc.Scan()
		cs[i].Stock, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		cs[i].Number, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(cs)
	var p int
label:
	for _, v := range cs {
		for i := 0; i < v.Stock; i++ {
			if p >= len(as) || v.Number < as[p] {
				break label
			}
			sum += int64(v.Number - as[p])
			p++
		}
	}
	fmt.Println(sum)
}
