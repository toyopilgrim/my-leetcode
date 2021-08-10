package reverse_int

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	strX := strconv.Itoa(abs(x))

	var reversedStrX string
	for i := len(strX) - 1; i >= 0; i-- {
		reversedStrX += string(strX[i])
	}

	if x < 0 {
		reversedStrX = "-" + reversedStrX
	}
	reversedX, _ := strconv.Atoi(reversedStrX)

	if reversedX < math.MinInt32 || reversedX > math.MaxInt32 {
		return 0
	}
	return reversedX
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
