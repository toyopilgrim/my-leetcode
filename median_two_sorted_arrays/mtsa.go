package median_two_sorted_arrays

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	cnt := len(nums)

	if cnt%2 != 0 {
		// Odd
		medianIdx := int(float64(cnt)/2+0.5) - 1
		return float64(nums[medianIdx])
	} else {
		// Even
		medianLeftIdx := cnt/2 - 1
		medianRightIdx := medianLeftIdx + 1
		median := float64(nums[medianLeftIdx]+nums[medianRightIdx]) / 2
		return median
	}
}
