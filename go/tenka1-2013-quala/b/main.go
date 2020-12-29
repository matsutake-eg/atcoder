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
	ans := 0
	for i := 0; i < n; i++ {
		v, w, x, y, z := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
		if v+w+x+y+z < 20 {
			ans++
		}
	}
	fmt.Println(ans)
}
