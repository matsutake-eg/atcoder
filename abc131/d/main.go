package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type work struct{ time, limit int }
type works []work

func (w works) Len() int      { return len(w) }
func (w works) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w works) Less(i, j int) bool {
	if w[i].limit == w[j].limit {
		return w[i].time < w[j].time
	}
	return w[i].limit < w[j].limit
}

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var ws works = make([]work, n)
	for i := range ws {
		sc.Scan()
		ws[i].time, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		ws[i].limit, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(ws)
	sum := int64(0)
	for _, w := range ws {
		sum += int64(w.time)
		if sum > int64(w.limit) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
