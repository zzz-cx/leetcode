package main

import "fmt"

var pairs = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if open, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

func main() {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"", true},
		{"(", false},
	}

	for _, tc := range tests {
		result := isValid(tc.s)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | s=%q => %v (expected %v)\n", status, tc.s, result, tc.expected)
	}
}
