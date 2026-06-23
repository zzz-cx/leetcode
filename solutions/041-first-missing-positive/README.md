# 缺失的第一个正数

> LeetCode 41 · [first-missing-positive](https://leetcode.cn/problems/first-missing-positive/)

## 题目

给定未排序的整数数组 `nums`，找出其中没有出现的最小正整数。要求 O(n) 时间、O(1) 额外空间。

## 题解思路与解析

- firstMissingPositive 将 [1,n] 内的数通过交换放到下标 i 上应满足 nums[i]==i+1，
- 再线性扫描第一个不符合的位置。时间 O(n)，额外空间 O(1)。

## 解答

### Golang

```go
// firstMissingPositive 将 [1,n] 内的数通过交换放到下标 i 上应满足 nums[i]==i+1，
// 再线性扫描第一个不符合的位置。时间 O(n)，额外空间 O(1)。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		v := nums[i]
		for v >= 1 && v <= n && nums[v-1] != v {
			nums[v-1], nums[i] = nums[i], nums[v-1]
			v = nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
```

### Python

```python
# firstMissingPositive 将 [1,n] 内的数通过交换放到下标 i 上应满足 nums[i]==i+1，
# 再线性扫描第一个不符合的位置。时间 O(n)，额外空间 O(1)。
from typing import List


class Solution:
    def first_missing_positive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            v = nums[i]
            while 1 <= v <= n and nums[v - 1] != v:
                nums[v - 1], nums[i] = nums[i], nums[v - 1]
                v = nums[i]
        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1
```
