# 找到字符串中所有字母异位词

> LeetCode 438 · [find-all-anagrams-in-a-string](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)

## 题目

给定字符串 `s` 和 `p`，找出 `s` 中所有 `p` 的字母异位词的起始下标。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func findAnagrams(s string, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}
```

### Python

```python
from typing import List


class Solution:
    def find_anagrams(self, s: str, p: str) -> List[int]:
        ans = []
        s_len, p_len = len(s), len(p)
        if s_len < p_len:
            return ans
        s_count = [0] * 26
        p_count = [0] * 26
        for i, ch in enumerate(p):
            s_count[ord(s[i]) - ord("a")] += 1
            p_count[ord(ch) - ord("a")] += 1
        if s_count == p_count:
            ans.append(0)

        for i, ch in enumerate(s[: s_len - p_len]):
            s_count[ord(ch) - ord("a")] -= 1
            s_count[ord(s[i + p_len]) - ord("a")] += 1
            if s_count == p_count:
                ans.append(i + 1)
        return ans
```
