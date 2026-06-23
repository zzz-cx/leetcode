# 划分字母区间

> LeetCode 763 · [partition-labels](https://leetcode.cn/problems/partition-labels/)

## 题目

给定字符串 `s`，将字符串划分为尽可能多的片段，使同一字母最多出现在一个片段中。

## 题解思路与解析

- 思路：贪心思想，记录每个字符最后出现的位置，然后遍历字符串，记录当前字符最后出现的位置，就是end，如果后续有字符的最后位置大于end了，就说明可以分割，小于就说明这个字符都在start到end的区间里

## 解答

### Golang

```go
// 思路：贪心思想，记录每个字符最后出现的位置，然后遍历字符串，记录当前字符最后出现的位置，就是end，如果后续有字符的最后位置大于end了，就说明可以分割，小于就说明这个字符都在start到end的区间里
func partitionLabels(s string) []int {
	var result []int
	lastIndex := make(map[rune]int)
	for i, ch := range s {
		lastIndex[ch] = i
	}
	start := 0
	end := 0
	for i, ch := range s {
		if lastIndex[ch] > end {
			end = lastIndex[ch]
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
```

### Python

```python
# 思路：贪心思想，记录每个字符最后出现的位置，然后遍历字符串，记录当前字符最后出现的位置，就是end
from typing import List


class Solution:
    def partition_labels(self, s: str) -> List[int]:
        result = []
        last_index = {}
        for i, ch in enumerate(s):
            last_index[ch] = i
        start = 0
        end = 0
        for i, ch in enumerate(s):
            if last_index[ch] > end:
                end = last_index[ch]
            if i == end:
                result.append(end - start + 1)
                start = end + 1
        return result
```
