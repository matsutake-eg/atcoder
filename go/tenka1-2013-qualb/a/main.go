package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	ss := make([]string, 15)
	for i := range ss {
		ss[i] = scanString()
	}
	sort.Strings(ss)
	fmt.Println(ss[6])
}
