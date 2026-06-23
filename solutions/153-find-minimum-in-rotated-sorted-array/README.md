# 寻找旋转排序数组中的最小值

> LeetCode 153 · [find-minimum-in-rotated-sorted-array](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/)

## 题目

已知旋转排序数组无重复，找出最小元素。

## 题解思路与解析

- 思路：二分缩区间，用 nums[mid] 与 nums[right] 比较判断最小值在左半还是右半

## 解答

### Golang

```go
// 思路：二分缩区间，用 nums[mid] 与 nums[right] 比较判断最小值在左半还是右半
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
```

### Python

```python
# 思路：二分缩区间，用 nums[mid] 与 nums[right] 比较判断最小值在左半还是右半
from typing import List


class Solution:
    def find_min(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        return nums[left]
```
