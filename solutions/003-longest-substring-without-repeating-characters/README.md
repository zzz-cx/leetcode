# 无重复字符的最长子串

> LeetCode 3 · [longest-substring-without-repeating-characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

## 题目

给定字符串 `s`，找出其中不含有重复字符的最长子串的长度。

## 题解思路与解析

- 检查 s[right] 是否在 [left, right-1] 窗口内重复
- 当前窗口长度 right-left+1，取最大值

## 解答

### Golang

```go
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
```

### Python

```python
class Solution:
    def length_of_longest_substring(self, s: str) -> int:
        n = len(s)
        if n == 0:
            return 0
        max_len = 1
        left = 0
        for right in range(1, n):
            for i in range(left, right):
                if s[i] == s[right]:
                    left = i + 1
                    break
            cur_len = right - left + 1
            if cur_len > max_len:
                max_len = cur_len
        return max_len
```
