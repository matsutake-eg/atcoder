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
	var ans int
	for i := 0; i < 5; i++ {
		x := scanInt()
		if x == 0 {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}
