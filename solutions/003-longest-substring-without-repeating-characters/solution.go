package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	maxLen := 1
	left := 0
	for right := 1; right < n; right++ {
		// 检查 s[right] 是否在 [left, right-1] 窗口内重复
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i + 1 // 重复则左边界移到重复字符的下一个
				break
			}
		}
		// 当前窗口长度 right-left+1，取最大值
		if curLen := right - left + 1; curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

func main() {
	tests := []struct{ s string; want int }{
		{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3},
	}
	for _, tc := range tests {
		got := lengthOfLongestSubstring(tc.s)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s | s=%q => %d (expected %d)\n", status, tc.s, got, tc.want)
	}
}
