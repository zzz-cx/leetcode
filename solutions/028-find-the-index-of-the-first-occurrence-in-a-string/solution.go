package main

import "fmt"

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

func main() {
	tests := []struct {
		haystack, needle string
		expected         int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "ll", 2},
		{"a", "a", 0},
	}

	for _, tc := range tests {
		result := strStr(tc.haystack, tc.needle)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | haystack=%q, needle=%q => %d (expected %d)\n", status, tc.haystack, tc.needle, result, tc.expected)
	}
}
