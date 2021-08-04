package substr_wo_repeat

var l int
var chars string

func LengthOfLongestSubstring(s string) int {
	var longestLength int
	chars = s

	l = len(s)
	for i := 0; i < l; i++ {
		subStr := []string{string(chars[i])}
		nextLength := len(extractSubStringWithoutRepeat(i, subStr))
		if longestLength < nextLength {
			longestLength = nextLength
		}
	}
	return longestLength
}

func extractSubStringWithoutRepeat(i int, subStr []string) []string {
	nextIndex := i + 1
	if nextIndex == l || contains(subStr, string(chars[nextIndex])) {
		return subStr
	}

	subStr = append(subStr, string(chars[nextIndex]))
	return extractSubStringWithoutRepeat(nextIndex, subStr)
}

func contains(s []string, n string) bool {
	for _, a := range s {
		if a == n {
			return true
		}
	}
	return false
}
