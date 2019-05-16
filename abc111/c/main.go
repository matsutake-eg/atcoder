package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Entry struct{ key, value int }
type List []Entry

func (l List) Len() int           { return len(l) }
func (l List) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l List) Less(i, j int) bool { return l[i].value > l[j].value }

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	om := make(map[int]int)
	em := make(map[int]int)
	var x int
	for i := 0; i < n/2; i++ {
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		om[x]++
		sc.Scan()
		x, _ = strconv.Atoi(sc.Text())
		em[x]++
	}

	var ol List = make([]Entry, len(om))
	index := 0
	for k, v := range om {
		ol[index].key, ol[index].value = k, v
		index++
	}
	sort.Sort(ol)

	var el List = make([]Entry, len(em))
	index = 0
	for k, v := range em {
		el[index].key, el[index].value = k, v
		index++
	}
	sort.Sort(el)

	if ol[0].key != el[0].key {
		fmt.Println(n - ol[0].value - el[0].value)
		return
	}

	if len(ol) == 1 {
		ol = append(ol, Entry{-1, 0})
	}
	if len(el) == 1 {
		el = append(el, Entry{-1, 0})
	}
	if v1, v2 := n-ol[0].value-el[1].value, n-ol[1].value-el[0].value; v1 < v2 {
		fmt.Println(v1)
	} else {
		fmt.Println(v2)
	}
}
