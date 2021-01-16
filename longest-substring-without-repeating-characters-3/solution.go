package main

import "math"

func LengthOfLongestSubstring(s string) int {
	l,longest := 0, 0;
	set := make(map[rune]int)

	for i, c := range s {
		if _, ok := set[c]; ok && set[c] >= l {
			longest = int(math.Max(float64(longest), float64(i - l)))
			l = set[c] + 1
		}
		set[c] = i
	}

	return int(math.Max(float64(longest), float64(len(s) - l)))
}