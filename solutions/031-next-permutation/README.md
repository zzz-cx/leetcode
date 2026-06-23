# 下一个排列

> LeetCode 31 · [next-permutation](https://leetcode.cn/problems/next-permutation/)

## 题目

实现获取下一个排列的函数：将数字重新排列成字典序更大的下一个排列。若不存在则按升序重排。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse2(nums, i+1, len(nums)-1)
}
func reverse2(nums []int, start int, end int) {
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
    def next_permutation(self, nums: List[int]) -> None:
        i = len(nums) - 2
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1
        if i >= 0:
            j = len(nums) - 1
            while j >= 0 and nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        self._reverse2(nums, i + 1, len(nums) - 1)

    def _reverse2(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
```
