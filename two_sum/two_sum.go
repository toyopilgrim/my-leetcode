package two_sum

func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for i2, v2 := range nums {
			if i != i2 {
				sum := v + v2
				if target == sum {
					return []int{i, i2}
				}
			}
		}
	}
	return nil
}
