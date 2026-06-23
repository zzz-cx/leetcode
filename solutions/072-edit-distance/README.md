# 编辑距离

> LeetCode 72 · [edit-distance](https://leetcode.cn/problems/edit-distance/)

## 题目

给定两个单词 `word1` 和 `word2`，返回将 `word1` 转换成 `word2` 所使用的最少操作数（插入、删除、替换）。

## 题解思路与解析

- 思路：二维 DP。dp[i][j] = word1 前 i 个字符变成 word2 前 j 个字符的最少操作数
- 相等则 dp[i][j]=dp[i-1][j-1]；否则取 删/增/替 三者最小 +1
- 删 word1[i-1]  插 word2[j-1]  替换

## 解答

### Golang

```go
// 思路：二维 DP。dp[i][j] = word1 前 i 个字符变成 word2 前 j 个字符的最少操作数
// 相等则 dp[i][j]=dp[i-1][j-1]；否则取 删/增/替 三者最小 +1
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删 i 个
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插 j 个
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				//              删 word1[i-1]  插 word2[j-1]  替换
			}
		}
	}
	return dp[m][n]
}
```

### Python

```python
# 思路：二维 DP。dp[i][j] = word1 前 i 个字符变成 word2 前 j 个字符的最少操作数
# 相等则 dp[i][j]=dp[i-1][j-1]；否则取 删/增/替 三者最小 +1
class Solution:
    def min_distance(self, word1: str, word2: str) -> int:
        m, n = len(word1), len(word2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m + 1):
            dp[i][0] = i
        for j in range(n + 1):
            dp[0][j] = j
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if word1[i - 1] == word2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = 1 + min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1])
        return dp[m][n]
```
