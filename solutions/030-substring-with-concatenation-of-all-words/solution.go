package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLen := len(words[0])
	wordCount := len(words)
	target := make(map[string]int, wordCount)
	for _, w := range words {
		target[w]++
	}

	result := []int{}

	for offset := 0; offset < wordLen; offset++ {
		left := offset
		used := 0
		seen := make(map[string]int)

		for right := offset; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			seen[word]++
			used++

			for seen[word] > target[word] {
				leftWord := s[left : left+wordLen]
				seen[leftWord]--
				if seen[leftWord] == 0 {
					delete(seen, leftWord)
				}
				left += wordLen
				used--
			}

			if used == wordCount {
				result = append(result, left)
				leftWord := s[left : left+wordLen]
				seen[leftWord]--
				if seen[leftWord] == 0 {
					delete(seen, leftWord)
				}
				left += wordLen
				used--
			}
		}
	}

	return result
}

func main() {
	tests := []struct {
		s        string
		words    []string
		expected []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	}

	for _, tc := range tests {
		result := findSubstring(tc.s, tc.words)
		if !sameSlice(result, tc.expected) {
			fmt.Printf("FAIL | s=%q, words=%v => %v (expected %v)\n", tc.s, tc.words, result, tc.expected)
			continue
		}
		fmt.Printf("PASS | s=%q, words=%v => %v (expected %v)\n", tc.s, tc.words, result, tc.expected)
	}
}

func sameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[int]int)
	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}
