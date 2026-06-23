# 寻找重复数

> LeetCode 287 · [find-the-duplicate-number](https://leetcode.cn/problems/find-the-duplicate-number/)

## 题目

给定 n+1 个整数，其中每个整数都在 1 到 n 之间，只有一个重复数字，找出它。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```

### Python

```python
from typing import List


class Solution:
    def find_duplicate(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            count = sum(1 for num in nums if num <= mid)
            if count > mid:
                right = mid
            else:
                left = mid + 1
        return left
```
