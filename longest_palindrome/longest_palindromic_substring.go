package longest_palindrome

import (
	"strings"
)

func longestPalindrome(s string) string {
	sList := strings.Split(s, "")
	// key = length, val = substring
	p := make(map[int]string)

	if s == reverseStr(s) {
		return s
	}

	var maxLength int
	if len(sList) > 0 {
		putLargerVal(p, 1, sList[0], &maxLength)
	}

	process(sList, p, &maxLength)

	return p[maxLength]
}

func process(sList []string, p map[int]string, ml *int) map[int]string {
	if len(sList) == 0 {
		return p
	}
	subStr := sList[0]
	isSingleCollection := true
	for i := 1; i <= len(sList)-1; i++ {
		// "aaaa"
		next := sList[i]
		subStr += next
		l := len(subStr)

		if isSingleCollection {
			if sList[0] == next {
				if i == l-1 {
					putLargerVal(p, l, subStr, ml)
				}
				continue
			} else {
				isSingleCollection = false
				continue
			}
		}

		if subStr[0] == subStr[l-1] && divideAndCompare(subStr, l) {
			reversed := reverseStr(subStr)
			if subStr == reversed {
				putLargerVal(p, l, subStr, ml)
			}
		}
	}
	return process(sList[1:], p, ml)
}

func divideAndCompare(str string, l int) bool {
	if l <= 2 {
		return true
	}
	// divide and compare
	if l%2 == 0 {
		left := l/2 - 1
		right := l / 2
		return compare(left, right, str)
	} else {
		mid := int(float64(l)/2+0.5) - 1
		return compare(mid-1, mid+1, str)
	}
}

func compare(leftIdx int, rightIdx int, s string) bool {
	left := s[leftIdx]
	right := s[rightIdx]
	res := left == right
	if res == false || leftIdx == 0 {
		return res
	}
	return compare(leftIdx-1, rightIdx+1, s)
}

func setMaxLength(val int, maxLength *int) {
	if *maxLength < val {
		*maxLength = val
	}
}

func putLargerVal(p map[int]string, l int, v string, maxLength *int) {
	if _, ok := p[l]; !ok {
		p[l] = v
		setMaxLength(l, maxLength)
	}
}

func reverseStr(s string) string {
	var rs string
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	return rs
}
