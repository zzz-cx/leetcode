package main

import "fmt"

// longestCommonPrefix 纵向扫描：逐列比较所有字符串的同一位置字符
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != ch {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
		{[]string{""}, ""},
	}

	for _, tc := range tests {
		result := longestCommonPrefix(tc.strs)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | strs=%v => %q (expected %q)\n", status, tc.strs, result, tc.expected)
	}
}
