package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	p := make([]int, n)
	for i := range p {
		p[i] = readInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		if p[i] == i+1 {
			ans++
			if i+1 < n && p[i+1] == i+2 {
				i++
			}
		}
	}
	fmt.Println(ans)
}
