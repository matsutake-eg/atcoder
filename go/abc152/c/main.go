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

	ans := 1
	t := readInt()
	for i := 0; i < n-1; i++ {
		p := readInt()
		if p < t {
			t = p
			ans++
		}
	}
	fmt.Println(ans)
}
