package main

import (
	"fmt"
	"math"
)

func MaxSubArray(nums []int) int {
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

func main() {
	fmt.Println(MaxSubArray([]int{-2,-1,-3,-4,-1,-2,-1,-5,-4}))
}