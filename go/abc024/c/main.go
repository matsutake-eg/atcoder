package main

import (
	"bufio"
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
	_, d, k := scanInt(), scanInt(), scanInt()
	l, r := make([]int, d), make([]int, d)
	for i := 0; i < d; i++ {
		l[i], r[i] = scanInt(), scanInt()
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < k; i++ {
		s, t := scanInt(), scanInt()
		fd := s < t
		for j := 0; j < d; j++ {
			if fd {
				if s >= l[j] && s < r[j] {
					s = r[j]
				}
				if s >= t {
					wr.WriteString(strconv.Itoa(j+1) + "\n")
					break
				}
			} else {
				if s > l[j] && s <= r[j] {
					s = l[j]
				}
				if s <= t {
					wr.WriteString(strconv.Itoa(j+1) + "\n")
					break
				}
			}
		}
	}
	wr.Flush()
}
