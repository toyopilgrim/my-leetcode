package zigza_conv

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	sList := strings.Split(s, "")
	rows := make([]string, numRows)

	diagonalRowNums := numRows - 2
	defaultThrottle := numRows + diagonalRowNums
	verticalThrottle := defaultThrottle

	for i, v := range sList {
		if i+1 <= numRows {
			rows[i] = v
			continue
		}

		if verticalThrottle == 0 {
			// Initialize throttle
			verticalThrottle = defaultThrottle
		}

		if verticalThrottle > numRows {
			diagonalRowIdx := verticalThrottle - numRows
			rows[diagonalRowIdx] += v
		} else {
			verticalRowIdx := numRows - verticalThrottle
			rows[verticalRowIdx] += v
		}
		verticalThrottle -= 1
	}

	var res string
	for _, v := range rows {
		res += v
	}
	return res
}
