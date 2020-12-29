package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	ans := 0
	ba, bx, xa := 0, 0, 0
	for i := 0; i < n; i++ {
		s := scanString()
		if s[0] == 'B' && s[len(s)-1] == 'A' {
			ba++
		} else if s[0] == 'B' {
			bx++
		} else if s[len(s)-1] == 'A' {
			xa++
		}
		for i := range s[:len(s)-1] {
			if s[i:i+2] == "AB" {
				ans++
			}
		}
	}
	if bx > 0 && ba > 0 {
		bx--
		ans++
	}
	if xa > 0 && ba > 0 {
		xa--
		ans++
	}
	if ba > 0 {
		ans += ba - 1
	}
	ans += min(bx, xa)
	fmt.Println(ans)
}
