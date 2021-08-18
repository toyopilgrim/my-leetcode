package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)

	reversed := reverse(xStr)

	return xStr == reversed
}

func reverse(s string) string {
	var rs string
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	return rs
}
