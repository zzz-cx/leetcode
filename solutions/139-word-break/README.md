# 单词拆分

> LeetCode 139 · [word-break](https://leetcode.cn/problems/word-break/)

## 题目

给定字符串 `s` 和字典 `wordDict`，判断 `s` 能否被拆分为若干字典单词。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word {
				dp[i] = dp[i] || dp[i-len(word)]
			}
		}
	}
	return dp[len(s)]
}
```

### Python

```python
from typing import List


class Solution:
    def word_break(self, s: str, word_dict: List[str]) -> bool:
        dp = [False] * (len(s) + 1)
        dp[0] = True
        for i in range(1, len(s) + 1):
            for word in word_dict:
                if i >= len(word) and s[i - len(word) : i] == word:
                    dp[i] = dp[i] or dp[i - len(word)]
        return dp[len(s)]
```
