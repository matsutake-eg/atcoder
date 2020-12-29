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
	n := scanInt()
	type work struct{ t, l int }
	ws := make([]work, n)
	for i := range ws {
		ws[i].t, ws[i].l = scanInt(), scanInt()
	}

	sort.Slice(ws, func(i, j int) bool { return ws[i].l < ws[j].l })
	sum := 0
	for _, w := range ws {
		sum += w.t
		if sum > w.l {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
