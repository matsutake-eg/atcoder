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
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	n := scanInt()
	sum := 0
	type city struct {
		s string
		p int
	}
	cs := make([]city, n)
	for i := range cs {
		s, p := scanString(), scanInt()
		cs[i].s, cs[i].p = s, p
		sum += p
	}

	sort.Slice(cs, func(i, j int) bool { return cs[i].p > cs[j].p })
	if cs[0].p > sum/2 {
		fmt.Println(cs[0].s)
	} else {
		fmt.Println("atcoder")
	}
}
