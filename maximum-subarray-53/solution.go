package maximum_subarray_53

import (
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	curMax := 0
	totMax := math.MinInt32

	for _, n := range nums {
		if curMax < 0 {
			curMax = n
		} else {
			curMax += n
		}

		totMax = maxInt(totMax, curMax)
	}

	return totMax
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
