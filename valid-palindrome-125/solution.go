package valid_palindrome_125

import "unicode"

func isPalindrome(s string) bool {
	var chars []rune

	for _, c := range s {
		if unicode.IsLetter(c) {
			chars = append(chars, unicode.ToLower(c))
		} else if unicode.IsDigit(c) {
			chars = append(chars, c)
		}
	}

	l := len(chars) - 1
	if l > 0 {
		for i := 0; i <= l / 2; i++ {
			if chars[i] != chars[l - i] {
				return false
			}
		}
	}

	return true
}