package main

import "fmt"

type stack []string

func (s stack) Push(c string) stack {
	return append(s, c)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func IsValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

	charMap := map[string]string{"(": ")", "{": "}", "[": "]"}

	var stack stack
	for _, c := range s {
		if _, ok := charMap[string(c)]; ok{
			stack = stack.Push(string(c))
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if charMap[top] == string(c) {
				stack, _ = stack.Pop()
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(IsValid("(){}}{"))
}
