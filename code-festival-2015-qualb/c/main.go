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
	n, m := scanInt(), scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	sort.Ints(a)
	b := make([]int, m)
	for i := range b {
		b[i] = scanInt()
	}
	sort.Ints(b)

	p := 0
	for i := range b {
		for j := p; j < len(a); j++ {
			if b[i] <= a[j] {
				if i == len(b)-1 {
					fmt.Println("YES")
					return
				}
				p = j + 1
				break
			}
		}
	}
	fmt.Println("NO")
}
