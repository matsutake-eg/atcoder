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
	idx := 1
	ans := 0
	for i := 0; i < n; i++ {
		a := readInt()
		if a != idx {
			ans++
			continue
		}
		idx++
	}

	if idx == 1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(ans)
}
