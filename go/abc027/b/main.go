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
	a := make([]int, n)
	total := 0
	for i := range a {
		a[i] = scanInt()
		total += a[i]
	}

	if total%n != 0 {
		fmt.Println(-1)
		return
	}

	ave := total / n
	cnt := 0
	sum := 0
	ans := 0
	for _, v := range a {
		cnt++
		sum += v
		if sum%cnt == 0 && sum/cnt == ave {
			cnt = 0
			sum = 0
		} else {
			ans++
		}
	}
	fmt.Println(ans)
}
