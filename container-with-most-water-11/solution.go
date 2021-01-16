package main

func MaxArea(height []int) int {
	curMax, l := 0, 0
	r := len(height) - 1

	for l < r {
		curArea := minInt(height[l], height[r]) * (r - l)
		curMax = maxInt(curMax, curArea)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return curMax
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}