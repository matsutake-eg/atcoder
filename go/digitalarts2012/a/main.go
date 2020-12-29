package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	s, n := scanString(), scanInt()
	tm := make(map[int]map[string]bool)
	for i := 0; i < n; i++ {
		t := scanString()
		l := len(t)
		if _, ok := tm[l]; !ok {
			tm[l] = make(map[string]bool)
		}
		tm[l][t] = true
	}

	ss := strings.Split(s, " ")
	wr := bufio.NewWriter(os.Stdout)
	for i, v := range ss {
		l := len(v)
		if _, ok := tm[l]; ok {
			found := false
			for ns := range tm[l] {
				if found {
					break
				}
				for i := range ns {
					if ns[i] == v[i] || ns[i] == '*' {
						if i == l-1 {
							found = true
						}
						continue
					}
					break
				}
			}
			if found {
				for range v {
					wr.WriteString("*")
				}
			} else {
				wr.WriteString(v)
			}
		} else {
			wr.WriteString(v)
		}
		if i < len(ss)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
