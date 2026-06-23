# 分割回文串

> LeetCode 131 · [palindrome-partitioning](https://leetcode.cn/problems/palindrome-partitioning/)

## 题目

给定字符串 `s`，将 `s` 分割成若干子串使每个子串都是回文串，返回所有可能的分割方案。

## 题解思路与解析

- partition 分割回文串（LeetCode 131）
- 思路：回文 DP 预处理 + 回溯。f[i][j] 表示 s[i..j] 是否为回文，dfs 枚举所有合法切法
- f[i][j]：s[i..j] 是否为回文。先全部置 true，单字符与「中间为空」自然成立
- 从右往左填表，保证算 f[i][j] 时 f[i+1][j-1] 已就绪
- splits：当前回溯路径上已切出的回文片段（等同 path）
- 切到末尾，复制一份 splits 加入 ans，避免后续回溯改乱
- 尝试下一段 s[i..j]

## 解答

### Golang

```go
// partition 分割回文串（LeetCode 131）
// 思路：回文 DP 预处理 + 回溯。f[i][j] 表示 s[i..j] 是否为回文，dfs 枚举所有合法切法
func partition(s string) (ans [][]string) {
	n := len(s)

	// f[i][j]：s[i..j] 是否为回文。先全部置 true，单字符与「中间为空」自然成立
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	// 从右往左填表，保证算 f[i][j] 时 f[i+1][j-1] 已就绪
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	// splits：当前回溯路径上已切出的回文片段（等同 path）
	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			// 切到末尾，复制一份 splits 加入 ans，避免后续回溯改乱
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		// 尝试下一段 s[i..j]
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1]) // 选
				dfs(j + 1)                          // 探：下一段从 j+1 开始
				splits = splits[:len(splits)-1]     // 撤
			}
		}
	}
	dfs(0)
	return
}
```

### Python

```python
# partition 分割回文串（LeetCode 131）
# 思路：回文 DP 预处理 + 回溯。f[i][j] 表示 s[i..j] 是否为回文，dfs 枚举所有合法切法
from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        n = len(s)
        f = [[True] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                f[i][j] = s[i] == s[j] and f[i + 1][j - 1]

        ans: List[List[str]] = []
        splits: List[str] = []

        def dfs(i: int) -> None:
            if i == n:
                ans.append(splits[:])
                return
            for j in range(i, n):
                if f[i][j]:
                    splits.append(s[i : j + 1])
                    dfs(j + 1)
                    splits.pop()

        dfs(0)
        return ans
```
