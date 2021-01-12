package two_sum_1

// TwoSum, given an array of integers, returns indices of the two numbers such
// that they add up to a specific target.
func TwoSum(nums []int, target int) []int {
	sum := []int{0,0}
	complements := make(map[int]int)

	for i, n := range nums {
		c := target - n
		if j, ok := complements[c]; ok {
			sum = []int{j, i}
			break
		} else {
			complements[n] = i
		}
	}

	return sum
}