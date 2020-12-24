package main

import (
	"bufio"
	"fmt"
	"os"
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
	dm := make(map[int]int)
	for i := 0; i < n; i++ {
		d := scanInt()
		dm[d]++
	}
	m := scanInt()
	for i := 0; i < m; i++ {
		t := scanInt()
		if dm[t] == 0 {
			fmt.Println("NO")
			return
		}
		dm[t]--
	}
	fmt.Println("YES")
}
