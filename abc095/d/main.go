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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n := readInt()
	a := make([]int, n)
	mx := 0
	for i := range a {
		a[i] = readInt()
		if a[i] > mx {
			mx = a[i]
		}
	}

	var (
		dhf = mx
		hf  int
	)
	for _, v := range a {
		if v == mx {
			continue
		}

		if d := abs(v - mx/2); d < dhf {
			dhf = d
			hf = v
		}
	}
	fmt.Println(mx, hf)
}
