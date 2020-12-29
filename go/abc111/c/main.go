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
	var n int
	fmt.Scan(&n)

	om := make(map[int]int, n/2)
	em := make(map[int]int, n/2)
	for i := 0; i < n/2; i++ {
		xo := scanInt()
		om[xo]++
		xe := scanInt()
		em[xe]++
	}

	type entry struct{ k, v int }
	ol := make([]entry, len(om))
	index := 0
	for k, v := range om {
		ol[index].k, ol[index].v = k, v
		index++
	}
	sort.Slice(ol, func(i, j int) bool { return ol[i].v > ol[j].v })

	el := make([]entry, len(em))
	index = 0
	for k, v := range em {
		el[index].k, el[index].v = k, v
		index++
	}
	sort.Slice(el, func(i, j int) bool { return el[i].v > el[j].v })

	if ol[0].k != el[0].k {
		fmt.Println(n - ol[0].v - el[0].v)
		return
	}

	if len(ol) == 1 {
		ol = append(ol, entry{-1, 0})
	}
	if len(el) == 1 {
		el = append(el, entry{-1, 0})
	}
	if v1, v2 := n-ol[0].v-el[1].v, n-ol[1].v-el[0].v; v1 < v2 {
		fmt.Println(v1)
	} else {
		fmt.Println(v2)
	}
}
