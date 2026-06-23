package main

import "fmt"

// 思路：中心扩展。每个位置作为回文中心，向两侧扩展；分奇长（i,i）和偶长（i,i+1）两种情况
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		// 循环结束时多扩出去一位，合法区间是 [left+1, right-1]
		if l := right - left - 1; l > maxLen {
			maxLen = l
			start = left + 1
		}
	}

	for i := range s {
		expand(i, i)   // 奇数长度，如 "aba"
		expand(i, i+1) // 偶数长度，如 "bb"
	}
	return s[start : start+maxLen]
}

func main() {
	got := longestPalindrome("babad")
	status := "PASS"
	if got != "bab" && got != "aba" {
		status = "FAIL"
	}
	fmt.Printf("%s | s=%q => %q\n", status, "babad", got)
}
