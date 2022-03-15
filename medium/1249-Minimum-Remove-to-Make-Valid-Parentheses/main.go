package main

import "strings"

func minRemoveToMakeValid(s string) string {
	numberOfCloses := 0
	for _, r := range s {
		if r == ')' {
			numberOfCloses++
		}
	}
	var sb strings.Builder
	n1, n2 := 0, 0
	for _, r := range s {
		switch r {
		case '(':
			if n1 < numberOfCloses {
				sb.WriteRune(r)
				n1++
			}
		case ')':
			if n1 > n2 {
				sb.WriteRune(r)
				n2++
			} else {
				numberOfCloses--
			}
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
