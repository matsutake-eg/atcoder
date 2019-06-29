package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	d := make([]int, n)
	for i := range d {
		sc.Scan()
		d[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(d)
	if d[n/2-1] == d[n/2] {
		fmt.Println(0)
		return
	}
	fmt.Println(d[n/2] - d[n/2-1])
}
