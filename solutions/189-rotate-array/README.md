# 旋转数组

> LeetCode 189 · [rotate-array](https://leetcode.cn/problems/rotate-array/)

## 题目

给定数组，将末尾 k 个元素整体移到前面，原地修改。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func rorate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
```

### Python

```python
from typing import List


class Solution:
    def rorate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        k = k % n
        self._reverse(nums, 0, n - 1)
        self._reverse(nums, 0, k - 1)
        self._reverse(nums, k, n - 1)

    def _reverse(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
```
