package valid_palindrome_125

import "unicode"

func isPalindrome(s string) bool {
	var chars []rune

	for _, c := range s {
		if unicode.IsLetter(c) {
			chars = append(chars, c)
		}
	}

	l := len(chars)
	if l > 1 {
		for i, c := range chars {
			if chars[i] != chars[l - i] {

			}
		}
	}

	return true
}