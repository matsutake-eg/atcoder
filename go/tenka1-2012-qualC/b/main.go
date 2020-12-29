package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scan(&s)

	xm := make(map[string]bool)
	xm["10"] = true
	xm["J"] = true
	xm["Q"] = true
	xm["K"] = true
	xm["A"] = true
	sm := make(map[string]bool)
	hm := make(map[string]bool)
	dm := make(map[string]bool)
	cm := make(map[string]bool)
	tm := make(map[string]bool)
L:
	for i := 0; i < len(s); {
		n := s[i+1 : i+2]
		if n == "1" {
			n = "10"
		}
		if xm[n] {
			m := s[i : i+1]
			switch m {
			case "S":
				sm[m+n] = true
				if len(sm) == 5 {
					tm = sm
					break L
				}
			case "H":
				hm[m+n] = true
				if len(hm) == 5 {
					tm = hm
					break L
				}
			case "D":
				dm[m+n] = true
				if len(dm) == 5 {
					tm = dm
					break L
				}
			case "C":
				cm[m+n] = true
				if len(cm) == 5 {
					tm = cm
					break L
				}
			}
		}
		if n == "10" {
			i += 3
		} else {
			i += 2
		}
	}

	miss := false
	wr := bufio.NewWriter(os.Stdout)
	am := make(map[string]bool)
	for i := 0; i < len(s); {
		t := s[i : i+2]
		if s[i+1:i+2] == "1" {
			t = s[i : i+3]
		}
		if tm[t] {
			am[t] = true
			if len(am) == 5 {
				if !miss {
					wr.WriteString("0")
				}
				wr.WriteString("\n")
				wr.Flush()
			}
		} else {
			miss = true
			wr.WriteString(t)
		}
		i += len(t)
	}
}
