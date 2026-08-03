package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := words[i]
		if mapped, ok := charToWord[ch]; ok {
			if mapped != word {
				return false
			}
		} else {
			if _, ok := wordToChar[word]; ok {
				return false
			}
			charToWord[ch] = word
			wordToChar[word] = ch
		}
	}
	return true
}

func main() {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"a", "dog", true},
	}

	for _, tc := range tests {
		result := wordPattern(tc.pattern, tc.s)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | pattern=%q, s=%q => %v (expected %v)\n", status, tc.pattern, tc.s, result, tc.expected)
	}
}
