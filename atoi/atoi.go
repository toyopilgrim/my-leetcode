package atoi

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	sList := strings.Split(s, "")

	var digitStr string
	var isNegative bool
	isInitial := true

	for _, v := range sList {
		if isInitial {
			if v == " " {
				continue
			} else if v == "-" {
				isNegative = true
				isInitial = false
				continue
			} else if v == "+" {
				isInitial = false
				continue
			}
		}

		_, err := strconv.Atoi(v)
		if err != nil {
			break
		}
		digitStr += v
		isInitial = false
	}

	if isNegative {
		digitStr = "-" + digitStr
	}

	res, _ := strconv.Atoi(digitStr)

	if res < 0 && res < math.MinInt32 {
		res = math.MinInt32
	}
	if res > 0 && res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}
