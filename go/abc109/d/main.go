package main

import (
	"bufio"
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

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return !isEven(x)
}

func main() {
	h, w := readInt(), readInt()
	a := make([][]int, h)
	for i := range a {
		a[i] = make([]int, w)
		for j := range a[i] {
			a[i][j] = readInt()
		}
	}

	type move struct{ y1, x1, y1d, x1d int }
	ans := make([]move, 0, h*w)
	cnt := 0
	for i := range a {
		if isEven(i) {
			for j := range a[i][:len(a[i])-1] {
				if isOdd(a[i][j]) {
					ans = append(ans, move{i + 1, j + 1, i + 1, j + 2})
					a[i][j+1]++
					cnt++
				}
			}
			if lai := len(a[i]) - 1; i != len(a)-1 && isOdd(a[i][lai]) {
				ans = append(ans, move{i + 1, lai + 1, i + 2, lai + 1})
				a[i+1][lai]++
				cnt++
			}
		} else {
			for j := len(a[i]) - 1; j > 0; j-- {
				if isOdd(a[i][j]) {
					ans = append(ans, move{i + 1, j + 1, i + 1, j})
					a[i][j-1]++
					cnt++
				}
			}
			if i != len(a)-1 && isOdd(a[i][0]) {
				ans = append(ans, move{i + 1, 1, i + 2, 1})
				a[i+1][0]++
				cnt++
			}
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	wr.WriteString(strconv.Itoa(cnt) + "\n")
	for _, m := range ans {
		wr.WriteString(strconv.Itoa(m.y1) + " " + strconv.Itoa(m.x1) + " " + strconv.Itoa(m.y1d) + " " + strconv.Itoa(m.x1d) + "\n")
	}
	wr.Flush()
}
