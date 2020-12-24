package main

import (
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	c := scanString()
	wr := bufio.NewWriter(os.Stdout)
	sp := false
	for _, r := range c {
		if sp && r != ' ' {
			sp = false
			wr.WriteString("," + string(r))
		} else if !sp && r == ' ' {
			sp = true
		} else if !sp && r != ' ' {
			wr.WriteString(string(r))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
