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
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		m[a]++
	}

	ans := 0
	for a, c := range m {
		if c >= a {
			ans += c - a
		} else {
			ans += c
		}
	}
	fmt.Println(ans)
}
