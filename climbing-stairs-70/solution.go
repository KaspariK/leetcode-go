package climbing_stairs_70

// DP solution
//func climbStairs(n int) int {
//	subProbs := make([]int, n + 1)
//	subProbs[0] = 1
//	subProbs[1] = 1
//
//	for i := 2; i <= n; i++ {
//		subProbs[i] = subProbs[i - 1] + subProbs[i - 2]
//	}
//
//	return subProbs[n]
//}

// Fibonacci solution
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a, b := 1, 2

	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}

	return b
}