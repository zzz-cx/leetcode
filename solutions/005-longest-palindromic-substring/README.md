# 最长回文子串

> LeetCode 5 · [longest-palindromic-substring](https://leetcode.cn/problems/longest-palindromic-substring/)

## 题目

给定字符串 `s`，找到 `s` 中最长的回文子串。

## 题解思路与解析

- 思路：中心扩展。每个位置作为回文中心，向两侧扩展；分奇长（i,i）和偶长（i,i+1）两种情况
- 循环结束时多扩出去一位，合法区间是 [left+1, right-1]

## 解答

### Golang

```go
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
```

### Python

```python
# 思路：中心扩展。每个位置作为回文中心，向两侧扩展；分奇长（i,i）和偶长（i,i+1）两种情况
class Solution:
    def longest_palindrome(self, s: str) -> str:
        if len(s) == 0:
            return ""
        start, max_len = 0, 1

        def expand(left: int, right: int) -> None:
            nonlocal start, max_len
            while left >= 0 and right < len(s) and s[left] == s[right]:
                left -= 1
                right += 1
            l = right - left - 1
            if l > max_len:
                max_len = l
                start = left + 1

        for i in range(len(s)):
            expand(i, i)
            expand(i, i + 1)
        return s[start : start + max_len]
```
