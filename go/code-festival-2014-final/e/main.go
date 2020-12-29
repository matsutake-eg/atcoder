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
	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x := scanInt()
		if i == 0 || r[len(r)-1] != x {
			r = append(r, x)
		}
	}

	if len(r) <= 2 {
		fmt.Println(0)
		return
	}

	ans := 1
	dr := true
	if r[1] < r[0] {
		dr = false
	}
	for i := 1; i < len(r); i++ {
		if dr && r[i] < r[i-1] {
			ans++
			dr = false
		} else if !dr && r[i] > r[i-1] {
			ans++
			dr = true
		}
	}
	if ans > 1 {
		fmt.Println(ans + 1)
	} else {
		fmt.Println(0)
	}
}
